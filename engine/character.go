package engine

import (
	"github.com/bunker-inspector/tba/domain"
)

func (e *Engine) NewCharacter(c *domain.Character) {
	e.GetCharacterRepo().Put(c)
}

func (e *Engine) DeleteCharacterByUserID(id int) {
	e.GetCharacterRepo().DeleteByUserID(id)
}

func (e *Engine) GetCharacterByUserID(id int) *domain.Character {
	return e.GetCharacterRepo().GetByUserID(id)
}
