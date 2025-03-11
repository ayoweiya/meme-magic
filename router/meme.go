package router

import (
	"github.com/gin-gonic/gin"
	"meme-magic/ai"
	"meme-magic/api"
)

func InitMemeRouter(r *gin.Engine) {
	memeGroup := r.Group("/memes")
	{
		memeGroup.POST("", api.CreateMemeHandler)       // 新增 Meme
		memeGroup.GET("/:id", api.GetMemeHandler)       // 取得單個 Meme
		memeGroup.PUT("", api.UpdateMemeHandler)        // 更新 Meme
		memeGroup.DELETE("/:id", api.DeleteMemeHandler) // 刪除 Meme
	}

	//支持不同 AI 生成模式
	r.POST("/generate/openai", api.GenerateMemeHandler(ai.GenerateMemeByOpenai)) //要收費就不要用了
	r.POST("/generate/replicate", api.GenerateMemeByReplicateHandler)            //要收費就不要用了
	r.POST("/generate/huggingFace", api.GenerateMemeByHuggingFaceHandler)        //要收費就不要用了
	r.POST("/generate/myLocalAI", api.GenerateMemeHandler(ai.GenerateMemeByMyLocalAI))

	telegramGroup := r.Group("/telegram")
	{
		telegramGroup.POST("/start", api.StartTelegramBotAPI)
		telegramGroup.POST("/stop", api.StopTelegramBotAPI)
	}
}
