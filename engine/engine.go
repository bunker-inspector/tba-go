package engine

type Engine struct {
	repo Repo
}

func NewEngine(r Repo) *Engine {
	return &Engine{repo: r}
}
