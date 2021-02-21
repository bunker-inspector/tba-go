package engine

import (
	"github.com/bunker-inspector/tba/domain"
)

func (e *Engine) NewCharacter(id int, c *domain.Character) {
	e.GetCharacterRepo().Put(id, c)
}

func (e *Engine) DeleteCharacterByUserID(id int) {
	e.GetCharacterRepo().DeleteByUserID(id)
}

func (e *Engine) GetCharacterByUserID(id int) *domain.Character {
	return e.GetCharacterRepo().GetByUserID(id)
}
