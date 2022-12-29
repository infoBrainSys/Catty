package openai

import (
	"GPT/config"
	"GPT/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cast"
)

// 聊天接口
const url = "https://api.openai.com/v1/completions"

// Chat 请求接口
func Chat(prompt string) (string, error) {
	responseChan := make(chan models.ResponseChat)
	errChan := make(chan error)

	go func() {
		// 实例 client, viper
		client := http.Client{}
		viper := config.Config()

		// 构建请求体
		reqChat := models.RequestChat{
			Model:       viper.GetString("RequestChat.Model"),
			Prompt:      prompt,
			MaxTokens:   viper.GetInt("RequestChat.MaxTokens"),
			Temperature: viper.GetFloat64("RequestChat.Temperature"),
			TopP:        viper.GetFloat64("RequestChat.TopP"),
			N:           viper.GetInt("RequestChat.N"),
		}

		body, err := json.Marshal(reqChat)
		if err != nil {
			errChan <- err
			log.Printf("Marshal err:%s", err)
			return
		}

		buf := bytes.NewBuffer(body)

		// 构建请求
		req, err := http.NewRequest("POST", url, buf)
		if err != nil {
			errChan <- err
			log.Printf("NewRequest err:%s", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+viper.GetString("RequestChat.Key"))

		// 发起请求并接收响应
		resp, err := client.Do(req)
		if err != nil {
			errChan <- err
			log.Printf("Do req err:%s", err)
			return
		}

		// 延时关闭请求体
		defer resp.Body.Close()

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
			log.Printf("ReadAll err:%s", err)
			return
		}

		// 反序列化结果到 models.ResponseChat
		var responseChat models.ResponseChat

		json.Unmarshal(bytes, &responseChat)

		// 响应体写入到 responseChan
		responseChan <- responseChat
	}()

	// 返回内容
	select {
	case text := <-responseChan:
		return cast.ToString(&text.Choices[0].Text), nil

	case err := <-errChan:
		return "", err
	}
}
