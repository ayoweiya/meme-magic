package api

import (
	"github.com/gin-gonic/gin"
	"meme-magic/service"
	"net/http"
)

var botInstance *service.TelegramBotService

func StartTelegramBotAPI(c *gin.Context) {
	if botInstance != nil && botInstance.IsRunning() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Telegram Bot 已經在運行"})
		return
	}

	botInstance = service.NewTelegramBotService()
	go botInstance.Start()

	c.JSON(http.StatusOK, gin.H{"message": "Telegram Bot 啟動成功"})
}

func StopTelegramBotAPI(c *gin.Context) {
	if botInstance == nil || !botInstance.IsRunning() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Telegram Bot 尚未啟動"})
		return
	}
	botInstance.Stop()
	botInstance = nil

	c.JSON(http.StatusOK, gin.H{"message": "Telegram Bot 已停止"})
}

func GetTelegramBotStatusAPI(c *gin.Context) {
	if botInstance == nil || !botInstance.IsRunning() {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Stop"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"status": "Start"})
}
