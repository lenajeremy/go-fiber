package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"primaryKey;type:uuid;unique" json:"id"`
	CreatedAt time.Time  `json:"createdAt" gorm:"created_at;autoCreateTime:nano;not null;default:now()"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"updated_at;autoUpdateTime:nano;not null;default:now()"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"deleted_at;index"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New()
	return nil
}
