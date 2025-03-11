package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"meme-magic/global"
	"meme-magic/initialize"
)

func RunWindowsServer() {
	r := gin.Default()    //初始化了一個 Gin 的路由引擎，它會自動加入一些預設
	initialize.Routers(r) // 註冊路由

	port := global.GVA_CONFIG.Port
	fmt.Println("Server is running on port:", port) // 顯示伺服器啟動資訊

	r.Run(":" + port)
}
