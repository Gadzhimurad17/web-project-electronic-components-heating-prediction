package config

type PostgresConfig struct {
	DbHost     string `env:"DB_HOST"`
	DbPort     int    `env:"DB_PORT"`
	DbUser     string `env:"DB_USER"`
	DbPassword int    `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME"`
}
