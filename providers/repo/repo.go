package repo

import (
	"log"

	"github.com/bunker-inspector/tba/config"
	"github.com/bunker-inspector/tba/engine"
	sqlite "github.com/bunker-inspector/tba/providers/repo/sqlite"
)

func NewRepoFactory(c *config.Config) *engine.RepoFactory {
	driver := *c.Storage()
	var factory engine.RepoFactory
	if driver == "sqlite" {
		factory = sqlite.NewSQLiteRepoFactory()
		return &factory
	}
	log.Fatal("STORAGE_DRIVER must be one of: sqlite")
	return nil
}
