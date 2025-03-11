package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"meme-magic/global"
	"meme-magic/model"
)

func Gorm() *gorm.DB {
	dbConfig := global.GVA_CONFIG.Database

	// 建立 MySQL DSN 連線字串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.Charset,
	)

	// 連接資料庫
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ 資料庫連接失敗: " + err.Error())
	}

	fmt.Println("✅ 資料庫連接成功！")

	err = db.AutoMigrate(&model.Meme{})
	if err != nil {
		panic("❌ 資料表創建失敗: " + err.Error())
	}
	fmt.Println("✅ 資料庫連接成功，Meme 資料表已建立！")

	return db
}
