package postgres

import (
	"context"

	"heating-prediction/internal/domain/models"

	"gorm.io/gorm"
)

type UserRepo struct{ db *gorm.DB }

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepo) FindUserByID(ctx context.Context, userID int) (*models.User, error) {
	var u models.User
	err := r.db.WithContext(ctx).First(&u, "user_id = ?", userID).Error
	return &u, err
}

func (r *UserRepo) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var u models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	return &u, err
}

// singletone
func (r *UserRepo) GetCurrentUserID() uint {
	return 1
}
