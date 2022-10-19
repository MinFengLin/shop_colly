package main

import (
	bot_service "bot"
	momo_colly "shop"
	"fmt"
	"strconv"

	/*
	 * env
	 */
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func tgbot_cmd(chatid *int64, token *string) {
	momo_data := momo_colly.Momo_parser()
	momo_info := "-\n"

	for ii := range momo_data.Momo {
		momo_info = momo_info + momo_data.Momo[ii].Item+"\n -> 目標價格："+momo_data.Momo[ii].Target_price + "\n 網址-(" +momo_data.Momo[ii].Url + ")" + "\n"
	}
	momo_info = momo_info + "-\n"
	bot_service.Telegram_reply_run(*chatid, *token, momo_info)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	chatID, _   := strconv.ParseInt(os.Getenv("chatID"), 10, 64)
	yourToken   := os.Getenv("yourToken")
	Env_Time, _ := strconv.ParseInt(os.Getenv("Timer_Minutes"), 10, 64)
	crontab_time := os.Getenv("Crontime")

	if len(crontab_time) > 0 {
		fmt.Printf("chatID: %d, yourToken: %s, Use crontab: %s \n", chatID, yourToken, crontab_time)
		c := cron.New()
		_, _= c.AddFunc(crontab_time, func() {
			momo_data := momo_colly.Momo_parser_data()
			bot_service.Telegram_bot_run(chatID, yourToken, momo_data)
		})
		momo_data := momo_colly.Momo_parser_data()
		bot_service.Telegram_bot_run(chatID, yourToken, momo_data)
		go tgbot_cmd(&chatID, &yourToken)
		c.Start()
		select{}
	} else {
		fmt.Printf("chatID: %d, yourToken: %s, How often to run: %d (Minutes)\n", chatID, yourToken, Env_Time)
		tChannel := time.NewTimer(time.Duration(Env_Time) * time.Minute)
		go tgbot_cmd(&chatID, &yourToken)
		for {
			momo_data := momo_colly.Momo_parser_data()
			bot_service.Telegram_bot_run(chatID, yourToken, momo_data)
			tChannel.Reset(time.Duration(Env_Time) * time.Minute)
			<-tChannel.C
		}
	}
}
