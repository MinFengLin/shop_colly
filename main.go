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
		c.AddFunc(crontab_time, func() {
			momo_data := momo_colly.Momo_parser()
			bot_service.Telegram_bot_run(chatID, yourToken, momo_data)
		})
		momo_data := momo_colly.Momo_parser()
		bot_service.Telegram_bot_run(chatID, yourToken, momo_data)
		c.Start()
		select{}
	} else {
		fmt.Printf("chatID: %d, yourToken: %s, How often to run: %d (Minutes)\n", chatID, yourToken, Env_Time)
		tChannel := time.NewTimer(time.Duration(Env_Time) * time.Minute)
		for {
			momo_data := momo_colly.Momo_parser()
			bot_service.Telegram_bot_run(chatID, yourToken, momo_data)
			tChannel.Reset(time.Duration(Env_Time) * time.Minute)
			<-tChannel.C
		}
	}
}
