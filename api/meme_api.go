package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"meme-magic/ai"
	"meme-magic/model"
	"meme-magic/service"
	"net/http"
	"strconv"
)

func CreateMemeHandler(c *gin.Context) {
	var meme model.Meme
	if err := c.ShouldBind(&meme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := service.CreateMeme(&meme); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message:": "新增成功", "meme": meme})
}

func GetMemeHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	meme, err := service.GetMemeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "找不到迷因"})
		return
	}
	c.JSON(http.StatusOK, meme)
}

func UpdateMemeHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var meme model.Meme
	if err := c.ShouldBindJSON(&meme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "請提供正確的 JSON 格式"})
		return
	}

	meme.ID = uint(id)
	if err := service.UpdateMeme(&meme); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失敗"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功", "meme": meme})
}

func DeleteMemeHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的 ID"})
		return
	}

	if err := service.DeleteMeme(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "刪除失敗"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "刪除成功"})
}

/*
*
共用版
*/
func GenerateMemeHandler(generatorFunc func(string) (string, error)) gin.HandlerFunc { //用閉包（Closure）回傳 gin.HandlerFunc
	return func(c *gin.Context) {
		var req struct {
			Prompt string `json:"prompt"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "請提供 prompt"})
			return
		}

		url, err := generatorFunc(req.Prompt)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "生成圖片失敗"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"image_url": url})
	}
}

func GenerateMemeByReplicateHandler(c *gin.Context) {
	var req struct {
		Prompt string `json:"prompt"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "請提供 prompt"})
		return
	}

	url, err := ai.GenerateMemeByReplicate(req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成圖片失敗"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_url": url})
}

func GenerateMemeByHuggingFaceHandler(c *gin.Context) {
	var req struct {
		Prompt string `json:"prompt"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "請提供 prompt"})
		return
	}

	url, err := ai.GenerateMemeByHuggingFace(req.Prompt)
	if err != nil {
		log.Println(err) // 顯示錯誤
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成圖片失敗"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_url": url})
}
