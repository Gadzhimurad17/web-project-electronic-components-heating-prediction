package postgres

import (
	"context"

	"heating-prediction/internal/domain/models"

	"gorm.io/gorm"
)

type ComponentRepo struct {
	db *gorm.DB
}

func NewComponentRepo(db *gorm.DB) *ComponentRepo {
	return &ComponentRepo{db: db}
}

func (r *ComponentRepo) CreateComponent(ctx context.Context, c *models.Component) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *ComponentRepo) ReadComponentByID(ctx context.Context, componentID uint16) (models.Component, error) {
	var receivedСomponent models.Component

	result := r.db.WithContext(ctx).
		Model(models.Component{}).
		Table("components").
		Select("*").
		Where("component_id = ?", componentID).
		Take(&receivedСomponent)
	if result.Error != nil {
		return models.Component{}, result.Error
	}

	return receivedСomponent, nil
}

func (r *ComponentRepo) ReadPublishedByID(ctx context.Context, componentID uint16) (models.Component, error) {
	var receivedСomponent models.Component

	result := r.db.WithContext(ctx).Model(models.Component{}).
		Table("components").Select("*").
		Where("component_id = ? AND component_status = ?", componentID, models.StatusPublished).
		Take(&receivedСomponent)
	if result.Error != nil {
		return models.Component{}, result.Error
	}
	return receivedСomponent, nil
}

func (r *ComponentRepo) ReadNextPublishedByID(ctx context.Context, componentID uint16) (models.Component, error) {
	var receivedСomponent models.Component

	result := r.db.WithContext(ctx).Model(models.Component{}).
		Table("components").Select("*").
		Where("component_id > ? AND component_status = ?", componentID, models.StatusPublished).
		Order("component_id ASC").
		First(&receivedСomponent).
		Take(&receivedСomponent)
	if result.Error != nil {
		return models.Component{}, result.Error
	}
	return receivedСomponent, nil
}

func (r *ComponentRepo) ReadComponentsList(ctx context.Context, minPower float64) ([]models.Component, error) {
	q := r.db.WithContext(ctx).
		Where("component_status = ?", models.StatusPublished)

	if minPower > 0 {
		q = q.Where("power_dissipation >= ?", minPower)
	}

	var components []models.Component
	err := q.Order("component_id ASC").Find(&components).Error
	if err != nil {
		return nil, err
	}
	return components, nil
}

func (r *ComponentRepo) ReadDraftByUser(ctx context.Context, userID uint) (models.Component, error) {
	var draft models.Component

	err := r.db.WithContext(ctx).
		Where("creator_id = ? AND component_status = ?", userID, models.StatusDraft).
		First(&draft).
		Error
	if err != nil {
		return models.Component{}, err
	}
	return draft, nil
}

func (r *ComponentRepo) UpdateComponent(ctx context.Context, component models.Component) error {
	result := r.db.WithContext(ctx).Model(models.Component{}).Where("component_id = ?", component.ComponentID).Updates(&models.Component{
		ComponentID:          component.ComponentID,
		ComponentName:        component.ComponentName,
		ComponentDescription: component.ComponentDescription,
		ComponentStatus:      component.ComponentStatus,
		ImageURL:             component.ImageURL,
		VideoURL:             component.VideoURL,
		ComponentType:        component.ComponentType,
		CreatedAt:            component.CreatedAt,
		CreatorID:            component.CreatorID,
		GeneratedAt:          component.GeneratedAt,
		PowerDissipation:     component.PowerDissipation,
		ThermalResistance:    component.ThermalResistance,
	})
	return result.Error
}

func (r *ComponentRepo) DeleteComponentSQL(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Exec(
		`UPDATE components SET component_status = $1 WHERE component_id = $2`,
		models.StatusDeleted, id,
	).Error
}
