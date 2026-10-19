package repo

import (
	"context"
	"heating-prediction/internal/domain/models"
)

type S3Repo interface {
	Upload(ctx context.Context, metadata *models.ContentMetadata) error
}
