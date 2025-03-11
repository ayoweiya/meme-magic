package ai

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"meme-magic/global"
)

// GenerateMeme 用 OpenAI DALL·E 生成迷因圖片
func GenerateMemeByOpenai(prompt string) (string, error) {
	client := openai.NewClient(global.GVA_CONFIG.OpenAIKey)

	resp, err := client.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Prompt:         prompt,
			N:              1,
			Size:           openai.CreateImageSize512x512, // 512x512 圖片
			ResponseFormat: openai.CreateImageResponseFormatURL,
		},
	)
	if err != nil {
		log.Println("❌ 生成圖片失敗:", err)
		return "", err
	}

	if len(resp.Data) == 0 {
		return "", fmt.Errorf("圖片生成失敗")
	}

	return resp.Data[0].URL, nil
}
