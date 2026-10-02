package gormrepo

import (
	"context"

	"heating-prediction/internal/domain/models"
	"heating-prediction/internal/domain/ports/repo"

	"gorm.io/gorm"
)

type ComponentRepo struct{ db *gorm.DB }

func NewComponentRepo(db *gorm.DB) repo.ComponentRepo {
	return &ComponentRepo{db: db}
}

func (r *ComponentRepo) CreateComponent(ctx context.Context, c *models.Component) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *ComponentRepo) GetComponentByID(ctx context.Context, id int) (*models.Component, error) {
	var c models.Component
	err := r.db.WithContext(ctx).
		Where("component_id = ? AND component_status <> ?", id, models.StatusDeleted).
		First(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ComponentRepo) GetPublishedComponents(ctx context.Context, minPower float64) ([]models.Component, error) {
	var list []models.Component
	q := r.db.WithContext(ctx).Where("component_status = ?", models.StatusPublished)
	if minPower > 0 {
		q = q.Where("power_dissipation >= ?", minPower)
	}
	err := q.Order("component_id").Find(&list).Error
	return list, err
}

func (r *ComponentRepo) GetDraftByCreator(ctx context.Context, creatorID int) (*models.Component, error) {
	var c models.Component
	err := r.db.WithContext(ctx).
		Where("creator_id = ? AND component_status = ?", creatorID, models.StatusDraft). // ← ИСПРАВЛЕНО
		First(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ComponentRepo) PublishComponent(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Model(&models.Component{}).
		Where("component_id = ?", id).
		Update("component_status", models.StatusPublished).Error
}

func (r *ComponentRepo) DeleteComponentSQL(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Exec(
		`UPDATE components SET component_status = $1 WHERE component_id = $2`,
		models.StatusDeleted, id,
	).Error
}


