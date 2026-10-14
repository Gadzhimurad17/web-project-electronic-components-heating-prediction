package minio

import (
	"fmt"
	"heating-prediction/internal/config"

	"github.com/minio/minio-go/v7"
)

func Connect(cfg *config.Config) (*minio.Client, error) {
	dsn, err := cfg.InitMinio()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize MinIO client: %w", err)
	}

	return dsn, nil
}
