package tgbot

import (
	"fmt"
	"log"
	"os"
	"time"

	tb "gopkg.in/telebot.v3"
)

type TgBot struct {
	bot *tb.Bot
}

type Weather interface {
	GetWeather(city string) string
}

func New() (*TgBot, error) {
	const op = "service.bot.New"

	bot, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TG_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &TgBot{bot: bot}, nil
}

func (t *TgBot) Start(weather Weather) {
	registerHandlers(t, weather)

	log.Println("Bot is working")

	t.bot.Start()
}

const (
	stateStart int = iota
	stateQuestion
	stateWeather
)

const (
	startCommand   = "/start"
	weatherCommand = "/weather"
	helpCommand    = "/help"
)

func registerHandlers(tbot *TgBot, weather Weather) {
	user := NewUserState()

	tbot.bot.Handle(tb.OnText, func(ctx tb.Context) error {
		return handleMessage(tbot, weather, user, ctx)
	})
}

func handleMessage(tbot *TgBot, weather Weather, user *User, ctx tb.Context) error {
	userId := ctx.Sender().ID
	message := ctx.Text()

	state, exists := user.GetState(userId)
	if !exists {
		user.SetState(userId, stateStart)
	}

	log.Printf("User State: %d, User ID: %d, Message: %s", state, userId, message)

	switch message {
	case helpCommand:
		return handleHelp(userId, message, ctx)
	default:
		switch state {
		case stateStart:
			return handleStart(userId, user, ctx)
		case stateQuestion:
			return handleQuestion(user, userId, ctx)
		case stateWeather:
			return handleWeather(weather, message, ctx)
		}
	}

	return nil
}

func handleStart(userId int64, user *User, ctx tb.Context) error {
	text := fmt.Sprintf("Hey %v, nice to meet you!\nUse 🌦weather button or command /weather to find out the 🌦weather in the place you are interested in!", ctx.Sender().Username)

	user.SetState(userId, stateQuestion)
	return ctx.Send(text)
}

func handleQuestion(user *User, userId int64, ctx tb.Context) error {
	user.SetState(userId, stateWeather)

	return ctx.Send("What city?\nGitHub: https://github.com/ZiganshinDev")

}

func handleWeather(weather Weather, message string, ctx tb.Context) error {
	req := weather.GetWeather(message)

	return ctx.Send(req)
}

func handleHelp(userId int64, message string, ctx tb.Context) error {
	return ctx.Send("This bot was created to help you find out about the 🌦weather in a place of your interest!\nYou can use 🌦weather button or command /weather for this")
}
