package models

type User struct {
	UserID       int    `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Email        string `gorm:"uniqueIndex;size:255" json:"email"`
	PasswordHash string `gorm:"size:255" json:"password_hash"`
	FirstName    string `gorm:"size:255" json:"first_name"`
	LastName     string `gorm:"size:255" json:"last_name"`

	Components []Component `gorm:"foreignKey:CreatorID;references:UserID" json:"-"`
	Likes      []Like      `gorm:"foreignKey:UserID;references:UserID" json:"-"`
}
