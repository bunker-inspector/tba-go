package rocksdb

import (
	rdb "github.com/bunker-inspector/gorocksdb"
	"github.com/bunker-inspector/tba/domain"
)

type characterRepo struct {
	DB *rdb.DB
}

func (r characterRepo) Get(c *domain.Character) *domain.Character {
	return c
}

func (r characterRepo) Put(c *domain.Character) {
}
