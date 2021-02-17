package telegram

import (
	//"fmt"
	// "log"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

func handleCharacterCommand(b *tb.Bot) func(*tb.Message) {
	return func (m *tb.Message) {
		subcommands := strings.Fields(m.Text)[1:]

		if subcommands[0] == "new" {
			newCharacter(b, m)
		} else {
			help(b, m)
		}
	}
}

func newCharacter(b *tb.Bot, m *tb.Message) {
	b.Send(m.Chat, "hello!")
}

func help(b *tb.Bot, m *tb.Message) {
	msg :=`
	/character subcommands
	-new [name]
	`
	b.Send(m.Chat, msg)
}
