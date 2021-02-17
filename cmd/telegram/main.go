package main

import (
	"log"
	"time"
	"os"

	"github.com/bunker-inspector/tba/adapters/telegram"
	"github.com/bunker-inspector/tba/cross"
	"github.com/bunker-inspector/tba/engine"
	"github.com/bunker-inspector/tba/providers/repo"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN not set.")
	}

	cross.InitConfig()
	config := cross.GetConfig()

	rf := repo.NewRepoFactory(config)

	engine.SetRepoFactory(rf)

	telegram.NewGame(token, 10 * time.Second).Start()
}
