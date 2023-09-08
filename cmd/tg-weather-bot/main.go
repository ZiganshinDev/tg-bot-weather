package main

import (
	"log"

	"github.com/ZiganshinDev/telebot/internal/service/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := bot.New()
	if err == nil {
		log.Println("Bot successfully init")
	}
	bot.Start()
}
