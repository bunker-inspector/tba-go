package rocksdb

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	rdb "github.com/bunker-inspector/gorocksdb"
	"github.com/bunker-inspector/tba/domain"
)

type characterRepo struct {
	DB      *rdb.DB
	RO      *rdb.ReadOptions
	WO      *rdb.WriteOptions
}

func key(id int) []byte {
	return []byte(fmt.Sprintf("character-%d", id))
}

func (r characterRepo) GetByPlayerID(id int) *domain.Character {
	key := key(id)

	data, err := r.DB.GetBytes(r.RO, []byte(key))

	if data == nil {
		log.Printf("No character found for User ID: %d\n", id)
		return nil
	}

	var c domain.Character
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatalf("decode error 1:", err)
	}
	return &c
}

func (r characterRepo) DeleteByPlayerID(id int) {
	key := key(id)
	r.DB.Delete(r.WO, key)
}


func (r characterRepo) Put(id int, c *domain.Character) {
	key := key(id)

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(*c)
	if err != nil {
		log.Printf("error: %+v\n", err)
	}
	r.DB.Put(r.WO, key, buffer.Bytes())
}
