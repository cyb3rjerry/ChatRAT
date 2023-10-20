package telegram

import (
	"errors"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	SlaveID string
	ChatID  int64
	Bot     *tgbotapi.BotAPI
}

func NewConn(token, id string) *TelegramBot {
	bot := new(TelegramBot)
	botInstance, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	bot.SlaveID = id
	bot.Bot = botInstance
	bot.ChatID = -1002121011166 //TODO: Handle this better, not a fan of hardcoding the channel

	return bot
}

func (tb TelegramBot) SendMessage() {
	message := tgbotapi.NewMessage(tb.ChatID, "")
	message.Text = "Test"
	fmt.Println("Sending message...")
	panic(errors.New("Not implemented"))
}

func (tb TelegramBot) ReceiveMessage() {
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := tb.Bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("Received '%s' from %s in group %#v", update.Message.Text, update.SentFrom().String(), update.FromChat())
	}
}

func (tb TelegramBot) SendMedia() {
	fmt.Println("Sending media...")
	panic(errors.New("Not implemented"))
}

func (tb TelegramBot) ReceiveMedia() {
	fmt.Println("Receiving media...")
	panic(errors.New("Not implemented"))
}
