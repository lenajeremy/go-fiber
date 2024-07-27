package models

import "github.com/google/uuid"

type Profile struct {
	BaseModel
	ProfilePicture string    `json:"profilePicture" gorm:"profile_picture;not null;default:https://images.unsplash.com/photo-1472099645785-5658abf4ff4e"`
	Gender         string    `json:"gender" gorm:"gender"`
	Email          string    `json:"email" gorm:"email;not null"`
	IsVerified     bool      `json:"isVerified" gorm:"is_verified;default:false;not null"`
	UserID         uuid.UUID `json:"userId" gorm:"user_id;type:uuid;not null;"`
}

type User2 struct {
	Profile Profile
}
