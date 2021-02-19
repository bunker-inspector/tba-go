package repo

import (
	"log"
	"github.com/bunker-inspector/tba/cross"
	"github.com/bunker-inspector/tba/engine"
	rdb "github.com/bunker-inspector/tba/providers/repo/rocksdb"
)

func NewRepoFactory(c *cross.Config) *engine.RepoFactory {
	driver := *c.Storage()

	var factory engine.RepoFactory
	if driver == "rocksdb" {
		factory = rdb.NewRocksDBRepoFactory()
		return &factory
	}

	log.Fatal("STORAGE_DRIVER must be one of: rocksdb")
	return nil
}
