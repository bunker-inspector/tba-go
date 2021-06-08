package engine

import (
	"github.com/bunker-inspector/tba/domain"
)

func (e *Engine) NewCharacter(c *domain.Character) {
	e.repo.SaveCharacter(c)
}

func (e *Engine) DeleteCharacterByUserID(id int) {
	e.repo.DeleteCharacterByUserID(id)
}

func (e *Engine) GetCharacterByUserID(id int) *domain.Character {
	return e.repo.GetCharacterByUserID(id)
}
