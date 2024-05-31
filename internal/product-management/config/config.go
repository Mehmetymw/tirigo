package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl       string
	RedisAddr   string
	RabbitMQUrl string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{
		DBUrl:       os.Getenv("DB_URL"),
		RedisAddr:   os.Getenv("REDIS_ADDR"),
		RabbitMQUrl: os.Getenv("RABBITMQ_URL"),
	}

	return config, nil
}
