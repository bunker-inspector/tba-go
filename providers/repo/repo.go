package repo

import (
	"log"

	"github.com/bunker-inspector/tba/config"
	"github.com/bunker-inspector/tba/engine"
	sqlite "github.com/bunker-inspector/tba/providers/repo/sqlite"
)

func NewRepo(c *config.Config) engine.Repo {
	driver := *c.Repo()
	if driver == "sqlite" {
		return sqlite.NewSQLiteRepo()
	}
	log.Fatal("REPO_DRIVER must be one of: sqlite")
	return nil
}
