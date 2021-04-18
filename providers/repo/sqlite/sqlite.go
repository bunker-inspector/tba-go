package sqlite

import (
	"github.com/bunker-inspector/tba/engine"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type sqliteRepoFactory struct {
	DB *gorm.DB
}

func NewSQLiteRepoFactory() engine.RepoFactory {
	db, err := gorm.Open(sqlite.Open("tba.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open SQLite DB at ./tba.db. Exiting...")
	}
	return sqliteRepoFactory{DB: db}
}

func (f sqliteRepoFactory) GetCharacterRepo() engine.CharacterRepo {
	return newCharacterRepo(f.DB)
}
