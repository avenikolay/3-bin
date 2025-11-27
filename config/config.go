package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("NO .env FILE")
	}
	res := os.Getenv("KEY")
	if res == "" {
		return nil, errors.New("NO KEY IN .env FILE")
	}
	return &Config{Key: res}, nil
}
