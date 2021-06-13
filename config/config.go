package config

import (
	dotenv "github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	repo string
}

var conf *Config

func InitConfig() {
	err := dotenv.Load()
	if err != nil {
		log.Print("No .env file found.")
	}

	repo := os.Getenv("REPO_DRIVER")
	if repo == "" {
		repo = "sqlite"
	}

	conf = &Config{
		repo: repo,
	}
}

func GetConfig() *Config {
	return conf
}

func (c *Config) Repo() *string {
	return &c.repo
}
