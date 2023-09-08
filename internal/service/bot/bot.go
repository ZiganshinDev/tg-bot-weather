package bot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ZiganshinDev/telebot/internal/service/weather"
	tb "gopkg.in/telebot.v3"
)

//TODO STATE

// type UserState struct {
// 	Step int
// 	// Другие поля, связанные с состоянием пользователя
// }

//var userStates map[int]*UserState

// func askForCity(c tb.Context, bot *tb.Bot, userID int) error {
//     state, exists := userStates[userID]

//     if !exists {
//         // Создайте новое состояние для нового пользователя
//         state = &UserState{}
//         userStates[userID] = state
//     }

//     switch state.Step {
//     case 0:
//         c.Send("What city?")
//         state.Step = 1

//         bot.Handle(tb.OnText, func(c tb.Context) error {
//             city := c.Text()
//             req := weather.ReflectWeather(weather.NewClient(city))
//             return c.Send(req)
//         })
//     case 1:
//         // Обработка ответа пользователя на город
//         // ...
//     }

//     return nil
// }

func New() (*tb.Bot, error) {
	const op = "service.bot.CreateBot"

	bot, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TG_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Printf("%s: %s", op, err)
		return nil, err
	}

	registerHandlers(bot)

	return bot, nil
}

func registerHandlers(bot *tb.Bot) {
	menu := &tb.ReplyMarkup{ResizeKeyboard: true}
	btnHelp := menu.Text("ℹ Help")
	btnWeather := menu.Text("🌦 Weather")
	menu.Reply(
		menu.Row(btnWeather),
		menu.Row(btnHelp),
	)

	bot.Handle("/start", handleStart(menu))
	bot.Handle("/weather", handleWeather(bot))
	bot.Handle("/help", handleHelp())
	bot.Handle(&btnWeather, handleWeather(bot))
	bot.Handle(&btnHelp, handleHelp())
}

func logMessage(userID int64, text string) {
	log.Printf("User ID: %d, Message: %s", userID, text)
}

func handleStart(menu *tb.ReplyMarkup) func(c tb.Context) error {
	return func(c tb.Context) error {
		userID := c.Sender().ID
		userText := c.Text()
		text := fmt.Sprintf("Hey %v, nice to meet you!\nUse 🌦weather button or command /weather to find out the 🌦weather in the place you are interested in!", c.Sender().Username)
		logMessage(userID, userText)
		return c.Send(text, menu)
	}
}

func handleWeather(bot *tb.Bot) func(c tb.Context) error {
	return func(c tb.Context) error {
		userID := c.Sender().ID
		userText := c.Text()
		if err := c.Send("What city?"); err != nil {
			logMessage(userID, fmt.Sprintf("Error sending message: %v", err))
			return err
		}

		logMessage(userID, userText)

		bot.Handle(tb.OnText, func(c tb.Context) error {
			city := c.Text()
			req := weather.ReflectWeather(weather.NewClient(city))
			logMessage(userID, fmt.Sprintf("Requested weather for city: %s", city))
			return c.Send(req)
		})

		return c.Send("GitHub: https://github.com/ZiganshinDev")
	}
}

func handleHelp() func(c tb.Context) error {
	return func(c tb.Context) error {
		userID := c.Sender().ID
		userText := c.Text()
		logMessage(userID, userText)
		return c.Send("This bot was created to help you find out about the 🌦weather in a place of your interest!\nYou can use 🌦weather button or command /weather for this")
	}
}
