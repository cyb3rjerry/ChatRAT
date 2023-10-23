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
func (tb DiscordBot) SendMessage() {
	fmt.Println("Sending message...")
	panic("Not implemented")

}

// TODO: remove testing stuff
func (tb DiscordBot) ReceiveMessage() {
	fmt.Println("Receiving message...")
	panic("Not implemented")
}

func (tb DiscordBot) SendMedia() {
	fmt.Println("Sending media...")
	panic("Not implemented")
}

func (tb DiscordBot) ReceiveMedia() {
	fmt.Println("Receiving media...")
	panic("Not implemented")
}
