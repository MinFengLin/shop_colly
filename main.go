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
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	chatID, _   := strconv.ParseInt(os.Getenv("chatID"), 10, 64)
	yourToken   := os.Getenv("yourToken")
	Env_Time, _ := strconv.ParseInt(os.Getenv("Timer_Minutes"), 10, 64)

	fmt.Printf("chatID: %d, yourToken: %s, Env_Time: %d (Seconds)\n", chatID, yourToken, Env_Time)
	tChannel := time.NewTimer(time.Duration(Env_Time) * time.Second)
	for {
		tChannel.Reset(time.Duration(Env_Time) * time.Second)
		<-tChannel.C
		momo_data := momo_colly.Momo_parser()
		bot_service.Telegram_bot_run(chatID, yourToken, momo_data)
	 }
}
