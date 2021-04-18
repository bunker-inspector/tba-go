package repo

import (
	"github.com/bunker-inspector/tba/cross"
	"github.com/bunker-inspector/tba/engine"
	sqlite "github.com/bunker-inspector/tba/providers/repo/sqlite"
	"log"
)

func NewRepoFactory(c *cross.Config) *engine.RepoFactory {
	driver := *c.Storage()
	var factory engine.RepoFactory
	if driver == "sqlite" {
		factory = sqlite.NewSQLiteRepoFactory()
		return &factory
	}
	log.Fatal("STORAGE_DRIVER must be one of: sqlite")
	return nil
}
