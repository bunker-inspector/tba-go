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

func buildCharacterCreationUI(c *domain.Character, uid int, points int, b *tb.Bot, e *engine.Engine) (string, *tb.ReplyMarkup) {
	tagFn := func (btnId string) string {
		return fmt.Sprintf("cc%s%d", btnId, uid)
	}

	// NOTE Not really sure what is valid but the callbacks
	// silently failed to route when hyphens were included.
	// Best to just use alphanumeric strings for now
	var (
		ui    = &tb.ReplyMarkup{}
		strDn = ui.Data("Str -", tagFn("strdn"))
		strUp = ui.Data("Str +", tagFn("strup"))
		conDn = ui.Data("Con -", tagFn("condn"))
		conUp = ui.Data("Con +", tagFn("conup"))
		dexDn = ui.Data("Dex -", tagFn("dexdn"))
		dexUp = ui.Data("Dex +", tagFn("dexup"))
		intDn = ui.Data("Int -", tagFn("intdn"))
		intUp = ui.Data("Int +", tagFn("intup"))
		wisDn = ui.Data("Wis -", tagFn("wisdn"))
		wisUp = ui.Data("Wis +", tagFn("wisup"))
		chaDn = ui.Data("Cha -", tagFn("chadn"))
		chaUp = ui.Data("Cha +", tagFn("chaup"))

		doneBtn = ui.Data("Done", tagFn("done"))

		rows []tb.Row
	)
	if c.Str > domain.ABILITY_MIN && c.Str < domain.ABILITY_INIT_MAX && points > 0 {
		rows = append(rows, ui.Row(strDn, strUp))
	} else if c.Str == domain.ABILITY_INIT_MAX || points == 0 {
		rows = append(rows, ui.Row(strDn))
	} else if c.Str == domain.ABILITY_MIN && points > 0 {
		rows = append(rows, ui.Row(strUp))
	}
	if c.Con > domain.ABILITY_MIN && c.Con < domain.ABILITY_INIT_MAX && points > 0 {
		rows = append(rows, ui.Row(conDn, conUp))
	} else if c.Con == domain.ABILITY_INIT_MAX || points == 0 {
		rows = append(rows, ui.Row(conDn))
	} else if c.Con == domain.ABILITY_MIN && points > 0 {
		rows = append(rows, ui.Row(conUp))
	}
	if c.Dex > domain.ABILITY_MIN && c.Dex < domain.ABILITY_INIT_MAX && points > 0 {
		rows = append(rows, ui.Row(dexDn, dexUp))
	} else if c.Dex == domain.ABILITY_INIT_MAX || points == 0 {
		rows = append(rows, ui.Row(dexDn))
	} else if c.Dex == domain.ABILITY_MIN && points > 0 {
		rows = append(rows, ui.Row(dexUp))
	}
	if c.Int > domain.ABILITY_MIN && c.Int < domain.ABILITY_INIT_MAX && points > 0 {
		rows = append(rows, ui.Row(intDn, intUp))
	} else if c.Int == domain.ABILITY_INIT_MAX || points == 0 {
		rows = append(rows, ui.Row(intDn))
	} else if c.Int == domain.ABILITY_MIN && points > 0 {
		rows = append(rows, ui.Row(intUp))
	}
	if c.Wis > domain.ABILITY_MIN && c.Wis < domain.ABILITY_INIT_MAX && points > 0 {
		rows = append(rows, ui.Row(wisDn, wisUp))
	} else if c.Wis == domain.ABILITY_INIT_MAX || points == 0 {
		rows = append(rows, ui.Row(wisDn))
	} else if c.Wis == domain.ABILITY_MIN && points > 0 {
		rows = append(rows, ui.Row(wisUp))
	}
	if c.Cha > domain.ABILITY_MIN && c.Cha < domain.ABILITY_INIT_MAX && points > 0 {
		rows = append(rows, ui.Row(chaDn, chaUp))
	} else if c.Cha == domain.ABILITY_INIT_MAX || points == 0 {
		rows = append(rows, ui.Row(chaDn))
	} else if c.Cha == domain.ABILITY_MIN && points > 0 {
		rows = append(rows, ui.Row(chaUp))
	}
	if points == 0 {
		rows = append(rows, ui.Row(doneBtn))
	}

	ui.Inline(rows...)

	b.Handle(&strUp, func(cb *tb.Callback) {
		b.Respond(cb, &tb.CallbackResponse{})
		label, ui := buildCharacterCreationUI(c, uid, points+1, b, e)
		b.Edit(cb.Message, label, ui)
	})
	b.Handle(&conDn, func(cb *tb.Callback) {
		c.Con--
		label, ui := buildCharacterCreationUI(c, uid, points+1, b, e)
		b.Edit(cb.Message, label, ui)
	})
	b.Handle(&dexDn, func(cb *tb.Callback) {
		c.Dex--
		label, ui := buildCharacterCreationUI(c, uid, points+1, b, e)
		b.Edit(cb.Message, label, ui)
	})
	b.Handle(&intDn, func(cb *tb.Callback) {
		c.Int--
		label, ui := buildCharacterCreationUI(c, uid, points+1, b, e)
		b.Edit(cb.Message, label, ui)
	})
	b.Handle(&wisDn, func(cb *tb.Callback) {
		c.Wis--
		label, ui := buildCharacterCreationUI(c, uid, points+1, b, e)
		b.Edit(cb.Message, label, ui)
	})
	b.Handle(&chaDn, func(cb *tb.Callback) {
		c.Cha--
		label, ui := buildCharacterCreationUI(c, uid, points+1, b, e)
		b.Edit(cb.Message, label, ui)
	})

	b.Handle(&strUp, func(cb *tb.Callback) {
		c.Str++
		label, ui := buildCharacterCreationUI(c, uid, points-1, b, e)
		b.Edit(cb.Message, label, ui)
	})
	b.Handle(&conUp, func(cb *tb.Callback) {
		c.Con++
		label, ui := buildCharacterCreationUI(c, uid, points-1, b, e)
		b.Edit(cb.Message, label, ui)
	})
	b.Handle(&dexUp, func(cb *tb.Callback) {
		c.Dex++
		label, ui := buildCharacterCreationUI(c, uid, points-1, b, e)
		b.Edit(cb.Message, label, ui)
	})
	b.Handle(&intUp, func(cb *tb.Callback) {
		c.Int++
		label, ui := buildCharacterCreationUI(c, uid, points-1, b, e)
		b.Edit(cb.Message, label, ui)
	})
	b.Handle(&wisUp, func(cb *tb.Callback) {
		c.Wis++
		label, ui := buildCharacterCreationUI(c, uid, points-1, b, e)
		b.Edit(cb.Message, label, ui)
	})
	b.Handle(&chaUp, func(cb *tb.Callback) {
		c.Cha++
		label, ui := buildCharacterCreationUI(c, uid, points-1, b, e)
		b.Edit(cb.Message, label, ui)
	})
	b.Handle(&doneBtn, func(cb *tb.Callback) {
		b.Edit(cb.Message, fmt.Sprintf("Welcome, O Brave %s", c.Name))

		// Were done, ship it
		e.NewCharacter(uid, c)
	})

	msg := "Name: %s\n" +
		"Points remaining: %d\n" +
		"Str: %d\n" +
		"Con: %d\n" +
		"Dex: %d\n" +
		"Int: %d\n" +
		"Wis: %d\n" +
		"Cha: %d\n"
	msg = fmt.Sprintf(msg,
		c.Name,
		points,
		c.Str,
		c.Con,
		c.Dex,
		c.Int,
		c.Wis,
		c.Cha,
	)

	return msg, ui
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
	log.Printf("name: %s\n", name)

	c := domain.BaseCharacter(name)

	label, ui := buildCharacterCreationUI(&c, m.Sender.ID, 6, b, e)

	b.Send(m.Chat, label, ui)
}

func showOwnCharacter(b *tb.Bot, m *tb.Message, e *engine.Engine) {
	c := e.GetCharacterByUserID(m.Sender.ID)

	if c == nil {
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
			c.Name,
			c.Level,
			c.Exp,
			c.Str,
			c.Con,
			c.Dex,
			c.Int,
			c.Wis,
			c.Cha,
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
