package rocksdb

import (
	rdb "github.com/bunker-inspector/gorocksdb"
	"github.com/bunker-inspector/tba/engine"
	"log"
)

type rocksDBRepoFactory struct {
	DB *rdb.DB
	RO *rdb.ReadOptions
	WO *rdb.WriteOptions
}

func NewRocksDBRepoFactory() engine.RepoFactory {
	bbto := rdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(rdb.NewLRUCache(3 << 30))
	opts := rdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	ro := rdb.NewDefaultReadOptions()
	wo := rdb.NewDefaultWriteOptions()
	db, err := rdb.OpenDb(opts, "tba.db")

	if err != nil {
		log.Fatalf("Failed to open RocksDB: %+v\n", err)
	}

	return rocksDBRepoFactory{
		DB: db,
		RO: ro,
		WO: wo,
	}
}

func (f rocksDBRepoFactory) GetCharacterRepo() engine.CharacterRepo {
	return characterRepo{
		DB: f.DB,
		RO: f.RO,
		WO: f.WO,
	}
}
