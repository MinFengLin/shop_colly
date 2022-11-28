package main

import (
	"fmt"
	"strconv"

	/*
	 * env
	 */
	"log"
	"os"
	"time"

	bot_service "github.com/MinFengLin/shop_colly/bot"
	momo "github.com/MinFengLin/shop_colly/crawler"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	chatID, _ := strconv.ParseInt(os.Getenv("chatID"), 10, 64)
	yourToken := os.Getenv("yourToken")
	Env_Time, _ := strconv.ParseInt(os.Getenv("Timer_Minutes"), 10, 64)
	crontab_time := os.Getenv("Crontime")

	if len(crontab_time) > 0 {
		fmt.Printf("chatID: %d, yourToken: %s, Use crontab: %s \n", chatID, yourToken, crontab_time)
		c := cron.New()
		_, _ = c.AddFunc(crontab_time, func() {
			momo_data := momo.Momo_parser_data()
			bot_service.Telegram_bot_run(chatID, yourToken, momo_data)
		})
		momo_data := momo.Momo_parser_data()
		bot_service.Telegram_bot_run(chatID, yourToken, momo_data)
		go bot_service.Telegram_reply_run(&chatID, &yourToken)
		c.Start()
		select {}
	} else {
		fmt.Printf("chatID: %d, yourToken: %s, How often to run: %d (Minutes)\n", chatID, yourToken, Env_Time)
		tChannel := time.NewTimer(time.Duration(Env_Time) * time.Minute)
		go bot_service.Telegram_reply_run(&chatID, &yourToken)
		for {
			momo_data := momo.Momo_parser_data()
			bot_service.Telegram_bot_run(chatID, yourToken, momo_data)
			tChannel.Reset(time.Duration(Env_Time) * time.Minute)
			<-tChannel.C
		}
	}
}
