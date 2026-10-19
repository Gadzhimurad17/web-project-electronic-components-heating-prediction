package service

import (
	"context"
	"heating-prediction/internal/domain/models"
	"heating-prediction/internal/infrastructure/dto"
	"io"
)

type ComponentService interface {
	PublishDraft(ctx context.Context,
		componentID uint16,
		status models.ServiceStatusEnum,
	) error
	GetDraft(ctx context.Context) (models.Component, error)
	GetFeedById(ctx context.Context, componentID uint16, next bool) (*models.Component, error)
	GetFiltredComponentsList(ctx context.Context, minPower float64) ([]models.Component, error)
	DeleteComponent(ctx context.Context, id uint16) error
	LikeComponent(ctx context.Context, componentID uint16, value int) error
	CreateComponent(ctx context.Context, req *dto.ComponentRequest, video io.Reader, image io.Reader) error
}
