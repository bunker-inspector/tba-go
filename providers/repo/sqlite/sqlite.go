package sqlite

import (
	"github.com/bunker-inspector/tba/engine"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type repo struct {
	*gorm.DB
}

func NewSQLiteRepo() engine.Repo {
	db, err := gorm.Open(sqlite.Open("tba.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open SQLite DB at ./tba.db. Exiting...")
	}
	sqliteRepo := repo{db}
	return &sqliteRepo
}
