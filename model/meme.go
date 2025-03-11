package model

import "gorm.io/gorm"

// Meme 代表迷因圖片的資料結構
type Meme struct {
	gorm.Model
	Title    string `gorm:"type:varchar(255);not null" json:"title"`     // 標題
	ImageURL string `gorm:"type:varchar(512);not null" json:"image_url"` // 圖片網址
	Likes    int    `gorm:"default:0" json:"likes"`                      // 按讚數
}
