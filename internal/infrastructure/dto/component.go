package dto

import (
	"heating-prediction/internal/domain/models"
	"mime/multipart"
)
type ComponentRequest struct {
    ComponentName        string                `form:"component_name" binding:"required"`
    ComponentDescription string                `form:"component_description"`
    ComponentType        models.ComponentType  `form:"component_type" binding:"required"`
    VideoFile            *multipart.FileHeader `form:"video_file" binding:"required"`
    ImageFile            *multipart.FileHeader `form:"image_file" binding:"required"`
    PowerDissipation     float64               `form:"power_dissipation" binding:"gte=0"`
    ThermalResistance    float64               `form:"thermal_resistance" binding:"gte=0"`
}
type ComponentResponse struct {
}
type LikeRequest struct {
	Value int // 0 или 1
}
