package models

import "github.com/google/uuid"

type User struct {
	BaseModel
	Username  string    `json:"username" gorm:"unique;not null"`
	FirstName string    `json:"firstName" gorm:"first_name"`
	LastName  string    `json:"lastName" gorm:"last_name"`
	Password  string    `json:"password" gorm:"password"`
	Profile   Profile   `json:"profile"`
	ProfileId uuid.UUID `json:"profileId" gorm:"foreignKey:ID; not null"`
	Todos     []Todo    `json:"todos" gorm:"foreignKey:UserID"`
}
