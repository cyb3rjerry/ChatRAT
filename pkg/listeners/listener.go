package listeners

import (
	"github.com/cyb3rjerry/ChatRAT/pkg/listeners/discord"
	"github.com/cyb3rjerry/ChatRAT/pkg/listeners/slack"
	"github.com/cyb3rjerry/ChatRAT/pkg/listeners/telegram"
)

type CommType int64

type ListenerConfig interface {
	telegram.TelegramConfig | discord.DiscordConfig | slack.SlackConfig
}

type Listener interface {
	SendMessage()
	ReceiveMessage()
	SendMedia()
	ReceiveMedia()
}

func NewListener[T ListenerConfig](config T) Listener {
	switch config := any(config).(type) { // https://github.com/golang/go/issues/49206
	case telegram.TelegramConfig:
		return telegram.NewConn(config)
	case discord.DiscordConfig:
		return discord.NewConn(config)
	case slack.SlackConfig:
		return slack.NewConn(config)
	}
	return nil
}
