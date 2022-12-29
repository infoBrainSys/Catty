package models

import "gorm.io/gorm"

// 存储发送的 prompt
type Prompt struct {
	*gorm.Model
	Prompt string `json:"prompt"`
}

// 回复的内容 ResponseText
type ResponseText struct {
	Text string `json:"text"`
}

// 请求结构体 RequestChat
type RequestChat struct {
	Model       string `json:"model"`
	Prompt      string `json:"prompt"`
	MaxTokens   int    `json:"max_tokens"`
	Temperature float64    `json:"temperature"`
	TopP        float64    `json:"top_p"`
	N           int    `json:"n"`
}

// 回复结构体 ResponseChat
type ResponseChat struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
