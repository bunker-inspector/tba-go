package engine

import (
	"github.com/bunker-inspector/tba/domain"
)

type Repo interface {
	DeleteCharacterByUserID(int)
	GetCharacterByUserID(int) *domain.Character
	SaveCharacter(*domain.Character)
}
