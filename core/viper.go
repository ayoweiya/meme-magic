package core

import (
	"fmt"
	"github.com/spf13/viper"
	"meme-magic/global"
)

// Viper 讀取 `config.yaml`，並初始化 `global.GVA_CONFIG`
func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config.yaml") // 指定設定檔

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("❌ 讀取 config.yaml 失敗: %s", err))
	}

	// 解析到 global.GVA_CONFIG
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		panic(fmt.Errorf("❌ 解析 config 失敗: %s", err))
	}

	fmt.Println("✅ 配置加載成功:", global.GVA_CONFIG)
	return v
}
