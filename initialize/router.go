package initialize

import (
	"github.com/gin-gonic/gin"
	"meme-magic/api"
	"meme-magic/router"
)

func Routers(r *gin.Engine) {
	r.GET("/hello", api.Hello)
	router.InitMemeRouter(r) // 註冊 Meme 路由
}
