package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	AppConfig
	RedisConfig
	PostgresConfig
	MinioConfig
}

func LoadFromEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("can't load .env: %v", err)
	}

	cfg := &Config{}

	if err := env.Parse(&cfg.AppConfig); err != nil {
		log.Fatalf("can't parse app config: %v", err)
	}

	if err := env.Parse(&cfg.PostgresConfig); err != nil {
		log.Fatalf("can't parse postgres config: %v", err)
	}

	if err := env.Parse(&cfg.MinioConfig); err != nil {
		log.Fatalf("can't parse minio config: %v", err)
	}

	return cfg
}
