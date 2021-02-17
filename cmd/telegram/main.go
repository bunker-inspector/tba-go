package main

import (
	"log"
	"time"
	"os"

	"github.com/bunker-inspector/tba/adapters/telegram"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN not set.")
	}

	telegram.NewGame(token, 10 * time.Second).Start()
}
