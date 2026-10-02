package gormrepo

import (
	"context"
	"heating-prediction/internal/domain/models"
	"heating-prediction/internal/domain/ports/repo"

	"gorm.io/gorm"
)

type likeRepo struct{ db *gorm.DB }

func NewLikeRepo(db *gorm.DB) repo.LikeRepo { return &likeRepo{db: db} }

func (r *likeRepo) GetLikesByComponentID(ctx context.Context, componentID int) ([]models.Like, error) {
	var list []models.Like
	err := r.db.WithContext(ctx).Where("component_id = ?", componentID).Find(&list).Error
	return list, err
}

func (r *likeRepo) CountByComponent(ctx context.Context, componentID int) (int, error) {
	var n int64
	err := r.db.WithContext(ctx).Model(&models.Like{}).
		Where("component_id = ?", componentID).
		Count(&n).Error
	return int(n), err
}

func (r *likeRepo) AddLike(ctx context.Context, userID, componentID int) error {
	return r.db.WithContext(ctx).Create(&models.Like{
		UserID:      userID,
		ComponentID: componentID,
	}).Error
}

func (r *likeRepo) RemoveLike(ctx context.Context, userID, componentID int) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND component_id = ?", userID, componentID).
		Delete(&models.Like{}).Error
}
