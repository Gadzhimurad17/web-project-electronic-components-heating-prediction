package models

import "time"

type ServiceStatusEnum string

const (
	StatusDraft     ServiceStatusEnum = "черновик"
	StatusPublished ServiceStatusEnum = "опубликован"
	StatusDeleted   ServiceStatusEnum = "удален"
)

type ComponentType string

const (
	ComponentResistor        ComponentType = "Резистор"
	ComponentCapacitor       ComponentType = "Конденсатор"
	ComponentTransistor      ComponentType = "Транзистор"
	ComponentDiode           ComponentType = "Диод"
	ComponentMicrocontroller ComponentType = "Микроконтроллер"
)

type Component struct {
	ComponentID          int               `json:"component_id"`
	ComponentName        string            `json:"component_name"`
	ComponentDescription string            `json:"component_description"`
	ComponentStatus      ServiceStatusEnum `json:"component_status"`
	ImageURL             string            `json:"image_url"`
	VideoURL             string            `json:"video_url"`
	ComponentType        ComponentType     `json:"component_type"`
	CreatedAt            time.Time         `json:"created_at"`
	CreatorID            int               `json:"creator_id"`
	GeneratedAt          time.Time         `json:"generated_at"`
	PowerDissipation     float64           `json:"power_dissipation"`
	ThermalResistance    float64           `json:"thermal_resistance"`
}
