package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6303766243:AAH6WPDeFqmS80hEhWbuDPhvnir0HzVpD2A")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	fmt.Println("fjei")

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
