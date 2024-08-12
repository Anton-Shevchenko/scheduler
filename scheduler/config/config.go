package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ServerPort string
	TGToken    string
	TGChatId   string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		ServerPort: os.Getenv("HTTP_PORT"),
		TGToken:    os.Getenv("TG_TOKEN"),
		TGChatId:   os.Getenv("TG_CHAT_ID"),
	}

	return config, nil
}
