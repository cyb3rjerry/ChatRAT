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

func NewConn(config SlackConfig) *SlackBot {
	// TODO
	return nil
}

// TODO: remove testing stuff
func (tb SlackBot) SendMessage() {
	fmt.Println("Sending message...")
	panic("Not implemented")

}

// TODO: remove testing stuff
func (tb SlackBot) ReceiveMessage() {
	fmt.Println("Receiving message...")
	panic("Not implemented")
}

func (tb SlackBot) SendMedia() {
	fmt.Println("Sending media...")
	panic("Not implemented")
}

func (tb SlackBot) ReceiveMedia() {
	fmt.Println("Receiving media...")
	panic("Not implemented")
}
