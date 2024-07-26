package models

type Todo struct {
	BaseModel
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"userId" gorm:"not null"`
	DueDate     string `json:"dueDate" gorm:"column:due_date"`
	IsCompleted bool   `json:"isCompleted" gorm:"column:is_completed"`
}
