package ai

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"meme-magic/global"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

func GenerateMemeByMyLocalAI(prompt string) (string, error) {
	encodedPrompt := url.QueryEscape(prompt)
	apiURL := fmt.Sprintf(global.GVA_CONFIG.LocalAIUrL+"?title=%s", encodedPrompt)
	log.Println(apiURL)

	// JSON 編碼
	req, _ := http.NewRequest("GET", apiURL, nil)

	// 設定 Header
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json")

	// 發送請求
	client := &http.Client{
		Timeout: 120 * time.Second, // 設定請求超時時間為 120 秒
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("❌ 請求發送失敗: %v", err)
	}

	defer resp.Body.Close()

	// 檢查 Content-Type 是否為圖片
	if resp.Header.Get("Content-Type") == "image/png" || resp.Header.Get("Content-Type") == "image/jpeg" {
		// 讀取圖片數據
		imgData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("❌ 讀取圖片失敗: %v", err)
		}

		// 產生圖片檔名
		fileName := fmt.Sprintf("meme_%d.png", time.Now().Unix())

		ensureImageDirExists()
		// 先存到本地
		localPath := filepath.Join("static", "images", fileName)

		err = ioutil.WriteFile(localPath, imgData, 0644)
		if err != nil {
			return "", fmt.Errorf("❌ 保存圖片失敗: %v", err)
		}

		return localPath, nil

		// 再上傳到 MinIO**
		//imageURL, err := utils.UploadToMinIO(fileName, imgData)
		//if err != nil {
		//	return localPath, fmt.Errorf("⚠️ 本地儲存成功，但上傳 MinIO 失敗: %v", err)
		//}
		//
		//// 返回MinIO URL
		//return imageURL, nil
	}

	// 如果不是圖片，嘗試解析 JSON
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("❌ 讀取回應失敗: %v", err)
	}

	// 解析 JSON
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("❌ 解析 JSON 失敗: %v，回應: %s", err, string(body))
	}

	return "", errors.New("生成圖片失敗")
}

func ensureImageDirExists() {
	err := os.MkdirAll("static/images", os.ModePerm)
	if err != nil {
		log.Fatal("❌ 建立資料夾失敗:", err)
	}
}
