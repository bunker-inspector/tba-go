package telegram

import (
	"fmt"
	"github.com/bunker-inspector/tba/domain"
	"github.com/bunker-inspector/tba/engine"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"strings"
)

func handleCharacterCommand(b *tb.Bot, e *engine.Engine) func(*tb.Message) {
	return func(m *tb.Message) {
		subcommands := strings.Fields(m.Text)[1:]

		if subcommands[0] == "delete" {
			deleteOwnCharacter(b, m, e)
		} else if subcommands[0] == "new" {
			newCharacter(b, m, e)
		} else if subcommands[0] == "me" {
			showOwnCharacter(b, m, e)
		} else {
			help(b, m)
		}
	}
}

func deleteOwnCharacter(b *tb.Bot, m *tb.Message, e *engine.Engine) {
	e.DeleteCharacterByUserID(m.Sender.ID)
	b.Send(m.Chat, "You character slot is empty.")
}

func newCharacter(b *tb.Bot, m *tb.Message, e *engine.Engine) {
	if c := e.GetCharacterByUserID(m.Sender.ID); c != nil {
		msg := "You have an existing character: %s\n" +
			"Please delete this character before creating a new one.\n" +
			"'/character delete'"
		b.Send(m.Chat, fmt.Sprintf(msg, c.Name))
		return
	}

	name := strings.Join(strings.Fields(m.Text)[2:], " ")
	log.Printf("name: %+v\n", name)

	character := domain.BaseCharacter(name)
	e.NewCharacter(m.Sender.ID, &character)

	b.Send(m.Chat, fmt.Sprintf("Welcome, O Brave %s", name))
}

func showOwnCharacter(b *tb.Bot, m *tb.Message, e *engine.Engine) {
	character := e.GetCharacterByUserID(m.Sender.ID)

	if character == nil {
		msg := "You have not created a character.\n" +
			"You can make one with 'character new [name]'\n"

		b.Send(m.Chat, msg)
	} else {
		msg := "Name: %s\n" +
			"Level: %d\n" +
			"Exp: %d\n\n" +
			"Str: %d\n" +
			"Con: %d\n" +
			"Dex: %d\n" +
			"Int: %d\n" +
			"Wis: %d\n" +
			"Cha: %d\n"
		msg = fmt.Sprintf(
			msg,
			character.Name,
			character.Level,
			character.Exp,
			character.Strength,
			character.Constitution,
			character.Dexterity,
			character.Intelligence,
			character.Wisdom,
			character.Charisma,
		)
		b.Send(m.Chat, msg)
	}
}

func help(b *tb.Bot, m *tb.Message) {
	msg := `
	/character subcommands
	-new [name]
	`
	b.Send(m.Chat, msg)
}
