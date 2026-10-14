package config

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioConfig struct {
	MinioEndpoint  string `env:"MINIO_ENDPOINT"`
	MinioAccesKey  string `env:"MINIO_ACCESS_KEY"`
	MinioSecretKey string `env:"MINIO_SECRET_KEY"`
	MinioBucket    string `env:"MINIO_BUCKET"`
	MinioUseSSL    bool   `env:"MINIO_USE_SSL"`
}

func (mcfg *MinioConfig) MINIO_DSN() string {
	return fmt.Sprintf(
		"endpoint=%s access_key=%s secret_key=%s bucket=%s ssl=%t",
		mcfg.MinioEndpoint,
		mcfg.MinioAccesKey,
		mcfg.MinioSecretKey,
		mcfg.MinioBucket,
		mcfg.MinioUseSSL,
	)
}

func (mcfg *MinioConfig) InitMinio() (*minio.Client, error) {
	return minio.New(mcfg.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(mcfg.MinioAccesKey, mcfg.MinioSecretKey, ""),
		Secure: mcfg.MinioUseSSL,
	})
}
