package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	AppConfig

	PostgresConfig
	MinioConfif
}

func LoadFromEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("can't load config settings from .env, error=%s", err)
	}

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("can't load config from env file, error=%s", err)
	}

	return cfg
}
