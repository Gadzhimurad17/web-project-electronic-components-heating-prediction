package repo

import (
	"context"
	"heating-prediction/internal/domain/models"
)

type LikeRepo interface {
	GetLikesByComponentID(ctx context.Context, componentID int) ([]models.Like, error)
	CountByComponent(ctx context.Context, componentID int) (int, error)
	AddLike(ctx context.Context, userID, componentID int) error
	RemoveLike(ctx context.Context, userID, componentID int) error
}
