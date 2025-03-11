package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"meme-magic/global"
	"net/http"
	"time"
)

// ReplicateAPI 請求結構
type ReplicateAPI struct {
	Version string `json:"version"`
	Input   struct {
		Prompt string `json:"prompt"`
	} `json:"input"`
}

// ReplicateResponse 回應結構
type ReplicateResponse struct {
	ID     string   `json:"id"`
	Output []string `json:"output,omitempty"`
	Status string   `json:"status"`
	Error  string   `json:"error,omitempty"`
}

// GenerateMemeByReplicate 使用 Replicate 生成圖片
func GenerateMemeByReplicate(prompt string) (string, error) {
	apiURL := "https://api.replicate.com/v1/predictions"

	// 設定請求內容
	reqBody := ReplicateAPI{
		Version: "db21e905a0e633f4287f3baf08414545b927477bfac123797c17f0fe10a6b7b8", // Stable Diffusion v1.5
	}
	reqBody.Input.Prompt = prompt

	jsonData, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))

	// 設定 Header
	req.Header.Set("Authorization", "Token "+global.GVA_CONFIG.ReplicateKey)
	req.Header.Set("Content-Type", "application/json")

	// 發送請求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 解析回應
	body, _ := ioutil.ReadAll(resp.Body)
	var result ReplicateResponse
	json.Unmarshal(body, &result)

	if result.Error != "" {
		return "", errors.New("Replicate API 錯誤: " + result.Error)
	}

	// 需要輪詢，等待生成完成
	predictionID := result.ID
	if predictionID == "" {
		return "", errors.New("未取得 prediction ID")
	}

	// 輪詢 API，等待圖片生成
	for i := 0; i < 10; i++ { // 最多輪詢 10 次
		time.Sleep(3 * time.Second) // 等待 3 秒

		url := "https://api.replicate.com/v1/predictions/" + predictionID
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Token "+global.GVA_CONFIG.ReplicateKey)

		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &result)

		// 如果生成成功，取得圖片 URL
		if result.Status == "succeeded" && len(result.Output) > 0 {
			return result.Output[0], nil
		} else if result.Status == "failed" {
			return "", errors.New("圖片生成失敗")
		}
	}

	return "", errors.New("圖片生成超時")
}
