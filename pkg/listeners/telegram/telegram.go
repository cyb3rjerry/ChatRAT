package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramConfig struct {
	ApiKey    string
	ChannelID int64
}

type TelegramBot struct {
	ChannelID int64
	Bot       *tgbotapi.BotAPI
}

func NewConn(config TelegramConfig) *TelegramBot {
	bot := new(TelegramBot)
	botInstance, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		panic(err)
	}

	bot.Bot = botInstance
	bot.ChannelID = config.ChannelID

	return bot
}

// TODO: remove testing stuff
func (tb TelegramBot) SendMessage() {
	fmt.Println("Sending message...")
	panic("Not implemented")

}

// TODO: remove testing stuff
func (tb TelegramBot) ReceiveMessage() {
	fmt.Println("Receiving message...")
	panic("Not implemented")
}

func (tb TelegramBot) SendMedia() {
	fmt.Println("Sending media...")
	panic("Not implemented")
}

func (tb TelegramBot) ReceiveMedia() {
	fmt.Println("Receiving media...")
	panic("Not implemented")
}
