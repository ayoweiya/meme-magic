package main

import (
	"meme-magic/core"
	"meme-magic/global"
	"meme-magic/initialize"
)

func main() {
	global.GVA_VP = core.Viper()      // 讀取 config.yaml
	global.GVA_DB = initialize.Gorm() // 連接資料庫
	core.RunWindowsServer()           // 啟動 Gin 伺服器
}
