package service

import (
	"context"
	"heating-prediction/internal/domain/models"
)

type AuthService interface {
	Register(ctx context.Context, email, password string) (*models.User, error)
}
