package minio

import (
	"context"
	"heating-prediction/internal/domain/models"

	"github.com/minio/minio-go/v7"
)

type minioFileRepo struct {
	mc     *minio.Client
	bucket string
}

func NewMinioFileRepo(mc *minio.Client, bucket string) *minioFileRepo {
	return &minioFileRepo{
		mc:     mc,
		bucket: bucket,
	}
}

func (r *minioFileRepo) Upload(ctx context.Context, metadata *models.ContentMetadata) error {
	opts := minio.PutObjectOptions{
		ContentType: metadata.MimeType,
	}
	_, err := r.mc.PutObject(ctx, r.bucket, metadata.URL, metadata.File, metadata.Size, opts)
	if err != nil {
		return models.ErrS3Operation
	}

	return nil
}
//func (r *minioFileRepo) Download(ctx context.Context)
