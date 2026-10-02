package models

type Like struct {
	LikeID      int `json:"like_id"`
	UserID      int `json:"user_id"`
	ComponentID int `json:"service_id"`
}
