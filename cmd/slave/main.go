package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/CoveoSec/chatrat/pkg/listeners"
	"github.com/CoveoSec/chatrat/pkg/slave"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	key := os.Getenv("SLAVE_API_KEY")

	commType := listeners.TELEGRAM

	slave := slave.NewSlave(key, commType)

	jsonSlave, _ := json.MarshalIndent(slave, "", "  ")

	fmt.Printf("%+v", string(jsonSlave))

}
