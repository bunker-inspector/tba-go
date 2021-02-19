package rocksdb

import (
	"log"

	rdb "github.com/bunker-inspector/gorocksdb"
	"github.com/bunker-inspector/tba/engine"
)

type rocksDBRepoFactory struct {
	DB *rdb.DB
}

func NewRocksDBRepoFactory() engine.RepoFactory {
	bbto := rdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(rdb.NewLRUCache(3 << 30))
	opts := rdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	db, err := rdb.OpenDb(opts, "tba.db")

	if err != nil {
		log.Fatalf("Failed to open RocksDB: %+v\n", err)
	}

	return rocksDBRepoFactory{DB: db}
}

func (f rocksDBRepoFactory) GetCharacterRepo() engine.CharacterRepo {
	return characterRepo{DB: f.DB}
}
