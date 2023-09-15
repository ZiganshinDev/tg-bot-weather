package main

import (
	"log"

	"github.com/ZiganshinDev/tg-bot-weather/services/tgbot/internal/client"
	"github.com/ZiganshinDev/tg-bot-weather/services/tgbot/internal/service/tgbot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	client, err := client.New()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbot.New()
	if err != nil {
		log.Fatal(err)
	}

	bot.Start(client)
}
