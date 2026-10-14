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
	ComponentID          int               `gorm:"primaryKey;autoIncrement" json:"component_id"`
	ComponentName        string            `gorm:"size:255;not null" json:"component_name"`
	ComponentDescription string            `gorm:"type:text" json:"component_description"`
	ComponentStatus      ServiceStatusEnum `gorm:"size:32;default:черновик;index" json:"component_status"`
	ImageURL             string            `gorm:"size:255" json:"image_url"`
	VideoURL             string            `gorm:"size:255" json:"video_url"`
	ComponentType        ComponentType     `gorm:"size:64" json:"component_type"`
	CreatedAt            time.Time         `gorm:"autoCreateTime" json:"created_at"`
	GeneratedAt          *time.Time        `json:"generated_at"`
	PowerDissipation     float64           `gorm:"not null" json:"power_dissipation"`
	ThermalResistance    float64           `gorm:"not null" json:"thermal_resistance"`

	CreatorID int  `gorm:"index" json:"creator_id"`
	Creator   User `gorm:"foreignKey:CreatorID;references:UserID;constraint:OnDelete:RESTRICT" json:"-"`
}
