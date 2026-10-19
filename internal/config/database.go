package config

import (
	"fmt"
	"net"
)

type PostgresConfig struct {
	DbHost     string `env:"DB_HOST"`
	DbPort     int    `env:"DB_PORT"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME"`
}

func (pcfg *PostgresConfig) POSTGRES_DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pcfg.DbHost,
		pcfg.DbPort,
		pcfg.DbUser,
		pcfg.DbPassword,
		pcfg.DbName,
	)
}

type RedisConfig struct {
	RedisHost     string `env:"REDIS_HOST"`
	RedisPort     string `env:"REDIS_PORT"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDatabase int    `env:"REDIS_DATABASE"`
}

func (r *RedisConfig) REDIS_DSN() string {
	return net.JoinHostPort(r.RedisHost, r.RedisPort)
}
