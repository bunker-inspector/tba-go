package cross

import (
	dotenv "github.com/joho/godotenv"
	"log"
	"os"
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
		storage = "sqlite"
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
