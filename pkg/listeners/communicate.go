package listeners

import "github.com/CoveoSec/chatrat/pkg/listeners/telegram"

type CommType int64

const (
	TELEGRAM CommType = iota
	DISCORD
	SLACK // TODO: Implement Slack
)

type Communicator interface {
	SendMessage()
	ReceiveMessage()
	SendMedia()
	ReceiveMedia()
}

func NewCommunicator(token, id string, botType CommType) Communicator {
	switch botType {
	case TELEGRAM:
		return telegram.NewConn(token, id)
	case DISCORD:
		// TODO: Implement Discord
		panic("Not implemented")
	case SLACK:
		// TODO: Implement Slack
		panic("Not implemented")
	}
	return nil
}
