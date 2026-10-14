package dto

import "heating-prediction/internal/domain/models"

type ComponentRequest struct {
    ComponentName        string
    ComponentDescription string
    ComponentType        models.ComponentType
    PowerDissipation     float64
    ThermalResistance    float64
}

type LikeRequest struct {
    Value int // 0 или 1
}
