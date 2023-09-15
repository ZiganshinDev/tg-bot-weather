package main

import (
	"flag"
	"log"

	"github.com/ZiganshinDev/tg-bot-weather/services/open-weather/internal/server"
	open_weather "github.com/ZiganshinDev/tg-bot-weather/services/open-weather/internal/service/open-weather"
	"github.com/joho/godotenv"
)

func main() {
	flag.Parse()

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	client, err := open_weather.New()
	if err != nil {
		log.Fatal(err)
	}

	s := server.New()
	s.Start(client)
}
