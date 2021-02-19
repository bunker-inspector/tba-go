package engine

import (
	"github.com/bunker-inspector/tba/domain"
)

type RepoFactory interface {
	GetCharacterRepo() CharacterRepo
}

type CharacterRepo interface {
	Get(*domain.Character) *domain.Character
	Put(*domain.Character)
}

var factory *RepoFactory

func SetRepoFactory(f *RepoFactory) {
	factory = f
}

func GetRepoFactory() *RepoFactory {
	return factory
}
