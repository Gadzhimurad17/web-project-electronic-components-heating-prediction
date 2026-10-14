package service

import (
	"context"
	"heating-prediction/internal/domain/models"
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
	// CreateComponent(
	// 	ctx context.Context,
	// 	req dto.CreateComponentRequest,
	// 	image io.Reader, imageName, imageCT string, imageSize int64,
	// 	video io.Reader, videoName, videoCT string, videoSize int64,
	// ) (*models.Component, error)
}
