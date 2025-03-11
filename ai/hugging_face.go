package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"meme-magic/global"
	"net/http"
	"time"
)

// HuggingFaceAPIRequest 用於請求 Hugging Face
type HuggingFaceAPIRequest struct {
	Inputs string `json:"inputs"`
}

// GenerateMemeByHuggingFace 透過 Hugging Face 生成圖片
func GenerateMemeByHuggingFace(prompt string) (string, error) {
	apiURL := global.GVA_CONFIG.HuggingFace.API_URL
	reqBody := HuggingFaceAPIRequest{Inputs: prompt}

	// JSON 編碼
	jsonData, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))

	// 設定 Header
	req.Header.Set("Authorization", "Bearer "+global.GVA_CONFIG.HuggingFace.API_Key)
	req.Header.Set("Content-Type", "application/json")

	// 發送請求
	client := &http.Client{
		Timeout: 120 * time.Second, // 設定為 120 秒
	}

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	// 解析回應
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("❌ 讀取回應失敗: %v", err)
	}

	// 解析 JSON
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("❌ 解析 JSON 失敗: %v，回應: %s", err, string(body))
	}

	// 檢查回應是否有錯誤
	if errMsg, ok := result["error"].(string); ok {
		return "", fmt.Errorf("❌ API 回應錯誤: %s", errMsg)
	}

	// 檢查回應是否有圖片 URL
	if url, ok := result["image"].(string); ok {
		return url, nil
	}

	return "", errors.New("生成圖片失敗")
}
