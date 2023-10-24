package discord

import (
	"fmt"
)

type DiscordConfig struct {
	// TODO
}

type DiscordBot struct {
	// TODO
}

func NewConn(config DiscordConfig) (*DiscordBot, error) {
	// TODO
	return nil, nil
}

// TODO: remove testing stuff
func (db DiscordBot) SendMessage(content string) error {
	fmt.Println("Sending message...")
	panic("Not implemented")

}

// TODO: remove testing stuff
func (db DiscordBot) GetMessages() ([]string, error) {
	fmt.Println("Receiving message...")
	panic("Not implemented")
}

func (db DiscordBot) SendMedia(text, fPath string) error {
	fmt.Println("Sending media...")
	panic("Not implemented")
}

func (db DiscordBot) GetLatestMedia(outPath string) error {
	fmt.Println("Receiving media...")
	panic("Not implemented")
}
