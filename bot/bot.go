package bot

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	momo "github.com/MinFengLin/shop_colly/crawler"
)

var (
	bot_r, bot_d *tgbotapi.BotAPI
)

func Tgbot_cmd(chatid *int64, token *string) {
	momo_data := momo.Momo_parser()
	momo_info := "-\n"

	for ii := range momo_data.Momo {
		momo_info = momo_info + momo_data.Momo[ii].Item+"\n -> 目標價格："+momo_data.Momo[ii].Target_price + "\n 網址-(" +momo_data.Momo[ii].Url + ")" + "\n"
	}
	momo_info = momo_info + "-\n"
	Telegram_reply_run(*chatid, *token, momo_info)
}

func sendMsg(msg string, chatID int64) {
	NewMsg := tgbotapi.NewMessage(chatID, msg)
	NewMsg.DisableWebPagePreview = true
	// NewMsg.ParseMode = tgbotapi.ModeHTML   //傳送html格式的訊息
	_, err := bot_d.Send(NewMsg)
	if err == nil {
		log.Printf("Send telegram message success")
	} else {
		log.Printf("Send telegram message error")
	}
}

func replyMsg(msg string, chatID int64) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates, _ := bot_r.GetUpdatesChan(updateConfig)
	for update_i := range updates {
		update := update_i
		go func() {
			cmd_text := update.Message.Text
			chatID := update.Message.Chat.ID
			replyMsg := tgbotapi.NewMessage(chatID, cmd_text)
			replyMsg.DisableWebPagePreview = true
			replyMsg.ReplyToMessageID = update.Message.MessageID
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "shop":
					replyMsg.Text = msg
				case "help":
					replyMsg.Text = "/shop  <-- to show all shop's items\n"
					replyMsg.Text = replyMsg.Text + "/shop_debug  <-- execute immediately shop crawler"
				case "shop_debug":
					replyMsg.Text = momo.Momo_parser_data()
				default:
					replyMsg.Text = ""
				}
			} else {
				replyMsg.Text = ""
			}
			if len(replyMsg.Text) > 0 {
				_, _ = bot_r.Send(replyMsg)
			}
		}()
	}
}

func Telegram_reply_run(chatID int64, yourToken string, msg string) {
	var err error
	bot_r, err = tgbotapi.NewBotAPI(yourToken)
	if err != nil {
		log.Fatal(err)
	}

	bot_r.Debug = false

	replyMsg(msg, chatID)
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
