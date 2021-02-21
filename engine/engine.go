package engine

type Engine struct {
	RepoFactory *RepoFactory
}

func NewEngine(f *RepoFactory) *Engine {
	return &Engine{RepoFactory: f}
}

func (e *Engine) GetRepoFactory() *RepoFactory {
	return e.RepoFactory
}

func (e *Engine) GetCharacterRepo() CharacterRepo {
	return (*e.GetRepoFactory()).GetCharacterRepo()
}
