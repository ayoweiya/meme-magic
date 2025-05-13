package main

import (
	"meme-magic/core"
	"meme-magic/global"
	"meme-magic/initialize"
	"meme-magic/service"
)

func main() {
	global.GVA_VP = core.Viper()      // 讀取 config.yaml
	global.GVA_DB = initialize.Gorm() // 連接資料庫

	// 啟動 Telegram Bot（用 goroutine）
	go func() {
		telegram := &service.TelegramBotService{}
		telegram.Start()
	}()

	// 啟動 Gin 伺服器（阻塞）
	core.RunWindowsServer() // 啟動 Gin 伺服器
}
