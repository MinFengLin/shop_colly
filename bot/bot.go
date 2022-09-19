package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	bot_d *tgbotapi.BotAPI
)

func sendMsg(msg string, chatID int64) {
	NewMsg := tgbotapi.NewMessage(chatID, msg)
	// NewMsg.ParseMode = tgbotapi.ModeHTML   //傳送html格式的訊息
	_, err := bot_d.Send(NewMsg)
	if err == nil {
		log.Printf("Send telegram message success")
	} else {
		log.Printf("Send telegram message error")
	}
}

func Telegram_bot_run(chatID int64, yourToken string, msg string) {
	var err error
	bot_d, err = tgbotapi.NewBotAPI(yourToken)
	if err != nil {
		log.Fatal(err)
	}

	bot_d.Debug = false

	sendMsg(msg, chatID)
}
