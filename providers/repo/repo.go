package repo

import (
	"log"

	"github.com/bunker-inspector/tba/cross"
	"github.com/bunker-inspector/tba/engine"
	sqlite "github.com/bunker-inspector/tba/providers/repo/sqlite"
)

func NewRepo(c *cross.Config) engine.Repo {
	driver := *c.Repo()
	if driver == "sqlite" {
		return sqlite.NewSQLiteRepo()
	}
	log.Fatal("REPO_DRIVER must be one of: sqlite")
	return nil
}
