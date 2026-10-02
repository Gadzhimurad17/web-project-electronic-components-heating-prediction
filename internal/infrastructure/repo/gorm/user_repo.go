package gormrepo

import (
	"context"

	"heating-prediction/internal/domain/models"
	"heating-prediction/internal/domain/ports/repo"

	"gorm.io/gorm"
)

type userRepo struct{ db *gorm.DB }

func NewUserRepo(db *gorm.DB) repo.UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepo) FindUserByID(ctx context.Context, userID int) (*models.User, error) {
	var u models.User
	err := r.db.WithContext(ctx).First(&u, "user_id = ?", userID).Error
	return &u, err
}

func (r *userRepo) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var u models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	return &u, err
}
