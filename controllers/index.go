package controllers

import (
	"GPT/openai"
	"log"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	prompt := ctx.Query("prompt")
	text, err := openai.Chat(prompt)
	if err != nil {
		log.Printf("openai.Chat err:%s", err)
		return
	}
	ctx.Writer.Write([]byte(text))
}
