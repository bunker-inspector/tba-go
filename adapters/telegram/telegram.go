package telegram

import (
	// "fmt"
	"github.com/bunker-inspector/tba/engine"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

func NewGame(token string, e *engine.Engine, timeout time.Duration) *tb.Bot {
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: timeout},
	})

	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/character", handleCharacterCommand(b, e))

	return b
}
