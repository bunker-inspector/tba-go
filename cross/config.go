package cross

import (
	"log"
	"os"
	dotenv "github.com/joho/godotenv"
)

type Config struct {
	storage string
}

var conf *Config

func InitConfig() {
	err := dotenv.Load()
	if err != nil {
		log.Print("No .env file found.")
	}

	storage := os.Getenv("STORAGE_DRIVER")
	if storage == "" {
		storage = "rocksdb"
	}

	conf = &Config{
		storage: storage,
	}
}

func GetConfig() *Config {
	return conf
}

func (c *Config) Storage() *string {
	return &(c.storage)
}
