package telegram

import (
	// "fmt"
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func NewGame(token string, timeout time.Duration) *tb.Bot {
	b, err := tb.NewBot(tb.Settings{
		Token: token ,
		Poller: &tb.LongPoller{Timeout: timeout},
	})

	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/character", handleCharacterCommand(b))

	return b
}
