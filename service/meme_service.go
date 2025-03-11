package service

import (
	"meme-magic/global"
	"meme-magic/model"
)

func CreateMeme(meme *model.Meme) error {
	return global.GVA_DB.Create(meme).Error
}

func GetMemeByID(id uint) (*model.Meme, error) {
	var meme model.Meme
	err := global.GVA_DB.First(&meme, id).Error
	return &meme, err
}

// 取得所有迷因
func GetAllMemes() ([]model.Meme, error) {
	var memes []model.Meme
	err := global.GVA_DB.Find(&memes).Error
	return memes, err
}

func UpdateMeme(meme *model.Meme) error {
	return global.GVA_DB.Save(meme).Error
}

// DeleteMeme 刪除 Meme
func DeleteMeme(id uint) error {
	return global.GVA_DB.Delete(&model.Meme{}, id).Error
}
