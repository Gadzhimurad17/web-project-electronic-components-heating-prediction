package repo

import (
	"context"
	"heating-prediction/internal/domain/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	FindUserByID(ctx context.Context, userID int) (*models.User, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
}
