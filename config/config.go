package config

// ServerConfig 定義伺服器相關的設定
type ServerConfig struct {
	Port             string         `mapstructure:"port" json:"port" yaml:"port"`
	Database         DatabaseConfig `mapstructure:"database" json:"database" yaml:"database"`
	OpenAIKey        string         `mapstructure:"openai_key"`
	ReplicateKey     string         `mapstructure:"replicate_key"`
	HuggingFace      HuggingFace    `mapstructure:"hugging_face"`
	LocalAIUrL       string         `mapstructure:"local_ai_url"`
	TelegramBotToken string         `mapstructure:"telegram_bot_token"`
}

type HuggingFace struct {
	API_URL string `mapstructure:"api_url"`
	API_Key string `mapstructure:"api_key"`
}

type DatabaseConfig struct {
	Type     string `mapstructure:"type" json:"type" yaml:"type"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Name     string `mapstructure:"name" json:"name" yaml:"name"`
	Charset  string `mapstructure:"charset" json:"charset" yaml:"charset"`
}
