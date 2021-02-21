package engine

import (
	"github.com/bunker-inspector/tba/domain"
)

type RepoFactory interface {
	GetCharacterRepo() CharacterRepo
}

type CharacterRepo interface {
	DeleteByUserID(int)
	GetByUserID(int) *domain.Character
	Put(int, *domain.Character)
}
