package slack

import (
	"fmt"
)

type SlackConfig struct {
	// TODO
}

type SlackBot struct {
	// TODO
}

func NewConn(config SlackConfig) (*SlackBot, error) {
	// TODO
	return nil, nil
}

// TODO: remove testing stuff
func (sb SlackBot) SendMessage(content string) error {
	fmt.Println("Sending message...")
	panic("Not implemented")

}

// TODO: remove testing stuff
func (sb SlackBot) GetMessages() ([]string, error) {
	fmt.Println("Receiving message...")
	panic("Not implemented")
}

func (sb SlackBot) SendMedia(text, fPath string) error {
	fmt.Println("Sending media...")
	panic("Not implemented")
}

func (sb SlackBot) GetLatestMedia(outPath string) error {
	fmt.Println("Receiving media...")
	panic("Not implemented")
}
