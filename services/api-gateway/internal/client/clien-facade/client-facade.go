package client_facade

import (
	"log"

	open_weather "github.com/ZiganshinDev/tg-bot-weather/services/api-gateway/internal/client/open-weather"
	yandex_weather "github.com/ZiganshinDev/tg-bot-weather/services/api-gateway/internal/client/yandex-weather"
)

type ClientFacade struct {
	yc *yandex_weather.Client
	oc *open_weather.Client
}

func New() *ClientFacade {
	yaClient, err := yandex_weather.New()
	if err != nil {
		log.Fatal("cannot init yandex_weather")
	}
	ocClient, err := open_weather.New()
	if err != nil {
		log.Fatal("cannot init open_weather")
	}

	return &ClientFacade{yc: yaClient, oc: ocClient}
}

func (c *ClientFacade) GetWeather(city string) string {
	return c.yc.GetWeather(c.oc.GetCityCoordinates(city))
}
