package service

import (
	"log"
	"meme-magic/ai"
	"meme-magic/global"
	"meme-magic/model"
	"sync"
	"time"

	"gopkg.in/tucnak/telebot.v2"
)

// TelegramBotService 用來管理 Bot 的啟動/停止
type TelegramBotService struct {
	bot     *telebot.Bot
	running bool
	mu      sync.Mutex
}

func NewTelegramBotService() *TelegramBotService {
	return &TelegramBotService{}
}

// 啟動 Telegram Bot
func (t *TelegramBotService) Start() {
	t.mu.Lock()
	defer t.mu.Unlock() // 確保函數結束時，資源會自動釋放

	if t.running {
		log.Println("⚠️ Telegram Bot 已經在運行")
		return
	}

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  global.GVA_CONFIG.TelegramBotToken, // 改成從 global 設定讀取
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal("❌ 啟動 Telegram Bot 失敗:", err)
	}

	t.bot = bot
	t.running = true

	// 設定指令
	bot.Handle("/meme", func(m *telebot.Message) {
		log.Println("收到 /meme 指令")
		prompt := "AI爆紅,吉卜力style, cute, anime style, vibrant colors, high-definition quality, natural look"
		imagePath, err := ai.GenerateMemeByMyLocalAI(prompt)
		if err != nil {
			log.Println(err)
			bot.Send(m.Chat, "❌ 生成迷因圖片失敗!")
			return
		}
		meme := model.Meme{
			Title:    prompt,
			ImageURL: imagePath,
		}
		if err := CreateMeme(&meme); err != nil {
			log.Println("❌ 儲存迷因到資料庫失敗:", err)
		}

		photo := &telebot.Photo{File: telebot.FromDisk(imagePath)}
		//Telegram 無法存取內網 IP 所以還是照原本的方式存在disk, 然後發佈到tg, 然後也上傳到minio
		//photo := &telebot.Photo{File: telebot.FromURL("http://192.168.68.57:9000/meme-magic/meme_1741725780.png")}
		bot.Send(m.Chat, photo)
	})

	// 開始監聽指令
	go bot.Start()
	log.Println("✅ Telegram Bot 已啟動")
}

// 停止 Telegram Bot
func (t *TelegramBotService) Stop() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.running {
		log.Println("⚠️ Telegram Bot 未在運行")
		return
	}

	t.bot.Stop()
	t.running = false
	log.Println("🛑 Telegram Bot 已停止")
}

// 檢查 Bot 是否在運行
func (t *TelegramBotService) IsRunning() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.running
}
