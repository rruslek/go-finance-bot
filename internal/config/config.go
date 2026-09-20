package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN"`
}

func Load() (*Config, error) {
	godotenv.Load()
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("TELEGRAM_BOT_TOKEN is not set")
	}

	return &Config{
		TelegramBotToken: token,
	}, nil
}
