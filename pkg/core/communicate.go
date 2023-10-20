package core

import "github.com/CoveoSec/chatrat/pkg/listeners/telegram"

type Slave interface {
	SendMessage()
	ReceiveMessage()
	SendMedia()
	ReceiveMedia()
}

func NewSlave(token, botType, id string) Slave {
	switch botType {
	case "telegram":
		return telegram.NewConn(token, id)
	}
	return nil
}
