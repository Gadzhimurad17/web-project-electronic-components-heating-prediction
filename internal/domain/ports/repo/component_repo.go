package repo

import (
	"context"
	"heating-prediction/internal/domain/models"
)

type ComponentRepo interface {
	CreateComponent(ctx context.Context, service *models.Component) error
	GetComponentByID(ctx context.Context, serviceID int) (*models.Component, error)
	GetPublishedComponents(ctx context.Context, minPower float64) ([]models.Component, error)
	GetDraftByCreator(ctx context.Context, creatorID int) (*models.Component, error)
	PublishComponent(ctx context.Context, id int) error
	DeleteComponentSQL(ctx context.Context, id int) error
}
