package bot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ZiganshinDev/telebot/pkg/weather"
	tb "gopkg.in/telebot.v3"
)

func CreateBot() {
	//creating bot
	newBot, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if newBot.Token == "" {
		log.Fatal("Env: Token must be set")
	}
	if err != nil {
		log.Fatal(err)
	}

	startBotHandels(newBot)

	newBot.Start()
}

func startBotHandels(newBot *tb.Bot) {
	menu := &tb.ReplyMarkup{ResizeKeyboard: true}

	btnHelp := menu.Text("ℹ Help")
	btnWeather := menu.Text("🌦 Weather")
	menu.Reply(
		menu.Row(btnWeather),
		menu.Row(btnHelp),
	)

	newBot.Handle("/start", func(c tb.Context) error {
		startMessage := fmt.Sprintf("Hey %v, nice to meet you!\nUse 🌦weather button or command /weather to find out the 🌦weather in the place you are interested in!", c.Sender().Username)

		return c.Send(startMessage, menu)
	})

	newBot.Handle("/weather", func(c tb.Context) error {
		c.Send("What city?")

		newBot.Handle(tb.OnText, func(c tb.Context) error {
			city := c.Text()
			req := weather.ReflectWeather(weather.NewClient(city))

			return c.Send(req)
		})

		return c.Send("GitHub: https://github.com/ZiganshinDev")
	})

	newBot.Handle(&btnWeather, func(c tb.Context) error {
		c.Send("What city?")

		newBot.Handle(tb.OnText, func(c tb.Context) error {
			city := c.Text()
			req := weather.ReflectWeather(weather.NewClient(city))

			return c.Send(req)
		})

		return c.Send("GitHub: https://github.com/ZiganshinDev")
	})

	newBot.Handle("/help", func(c tb.Context) error {
		return c.Send("This bot was created to help you find out about the 🌦weather in a place of your interest!\nYou can use 🌦weather button or command /weather for this")
	})

	newBot.Handle(&btnHelp, func(c tb.Context) error {
		return c.Send("This bot was created to help you find out about the 🌦weather in a place of your interest!\nYou can use 🌦weather button or command /weather for this")
	})
}
