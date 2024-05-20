package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Recommendation struct {
	ID         int64          `gorm:"primaryKey; autoIncrement"`
	UserID     uuid.UUID          `gorm:"index"`
	RecipeName string         `gorm:"type:varchar(255)"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
