package repo

import (
	"context"
	"heating-prediction/internal/domain/models"
)

type ComponentRepo interface {
	CreateComponent(ctx context.Context, service *models.Component) error

	ReadComponentByID(ctx context.Context, componentID uint16) (models.Component, error)
	ReadNextPublishedByID(ctx context.Context, componentID uint16) (models.Component, error)
	ReadPublishedByID(ctx context.Context, componentID uint16) (models.Component, error)
	ReadComponentsList(ctx context.Context, minPower float64) ([]models.Component, error)
	ReadDraftByUser(ctx context.Context, userID uint) (models.Component, error)

	UpdateComponent(ctx context.Context, component models.Component) error
	DeleteComponentSQL(ctx context.Context, id int) error
}
