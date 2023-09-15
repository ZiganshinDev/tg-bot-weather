package main

import (
	"flag"

	client_facade "github.com/ZiganshinDev/tg-bot-weather/services/api-gateway/internal/client/clien-facade"
	"github.com/ZiganshinDev/tg-bot-weather/services/api-gateway/internal/server"
)

func main() {
	flag.Parse()

	c := client_facade.New()

	s := server.New()

	s.Run(c)
}
