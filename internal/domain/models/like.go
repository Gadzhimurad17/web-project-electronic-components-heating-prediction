package models

type Like struct {
	LikeID int `gorm:"primaryKey;autoIncrement" json:"like_id"`

	UserID      int `gorm:"uniqueIndex:idx_user_component" json:"user_id"`
	ComponentID int `gorm:"uniqueIndex:idx_user_component" json:"component_id"`

	User      User      `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:RESTRICT" json:"-"`
	Component Component `gorm:"foreignKey:ComponentID;references:ComponentID;constraint:OnDelete:RESTRICT" json:"-"`
}
