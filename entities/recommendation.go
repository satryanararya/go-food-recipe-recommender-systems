package entities

import (
	"time"

	"gorm.io/gorm"
)

type Recommendation struct {
	ID         int64          `gorm:"primaryKey; autoIncrement"`
	UserID     int64          `gorm:"index"`
	RecipeName string         `gorm:"type:varchar(255)"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
