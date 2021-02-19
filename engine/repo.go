package engine

import (
	"github.com/bunker-inspector/tba/domain"
)

type RepoFactory interface {
	GetCharacterRepo() CharacterRepo
}

type CharacterRepo interface {
	DeleteByPlayerID(int)
	GetByPlayerID(int) *domain.Character
	Put(int, *domain.Character)
}

var factory *RepoFactory

func SetRepoFactory(f *RepoFactory) {
	factory = f
}

func GetRepoFactory() *RepoFactory {
	return factory
}
