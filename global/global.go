package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"meme-magic/config"
)

// 全域變數
var (
	GVA_CONFIG config.ServerConfig // GVA_CONFIG 全域變數，存放應用設定
	GVA_VP     *viper.Viper        // GVA_VP 存放 Viper 實例
	GVA_DB     *gorm.DB
)
