package config

import "fmt"

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
