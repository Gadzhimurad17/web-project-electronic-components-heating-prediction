package config

type MinioConfif struct {
	MinioEndpoint  int    `env:"MINIO_ENDPOINT"`
	MinioAccesKey  string `env:"MINIO_ACCESS_KEY"`
	MinioSecretKey string `env:"MINIO_SECRET_KEY"`
	MinioBucket    string `env:"MINIO_BUCKET"`
	MinioUseSsl    bool   `env:"MINIO_USE_SSL"`
}
