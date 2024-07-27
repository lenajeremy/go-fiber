package models

type User struct {
	BaseModel
	Username  string  `json:"username" gorm:"unique;not null"`
	FirstName string  `json:"firstName" gorm:"first_name"`
	LastName  string  `json:"lastName" gorm:"last_name"`
	Password  string  `json:"password" gorm:"password"`
	Profile   Profile `json:"profile" gorm:"foreignKey:UserID"`
	Todos     []Todo  `json:"todos" gorm:"foreignKey:UserID"`
}
