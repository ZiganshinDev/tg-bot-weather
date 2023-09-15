package main

import (
	"flag"
	"log"

	"github.com/ZiganshinDev/tg-bot-weather/services/yandex-weather/internal/server"
	yandex_weather "github.com/ZiganshinDev/tg-bot-weather/services/yandex-weather/internal/service/yandex-weather"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
	flag.Parse()

	client, err := yandex_weather.New()
	if err != nil {
		log.Fatalln("failed to init yandex-weather")
	}

	s := server.New()

	s.Start(client)
}
