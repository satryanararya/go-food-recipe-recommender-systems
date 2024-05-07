package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64  		`gorm:"primaryKey"`
	Username  string 		`gorm:"unique;not null"`
	Email     string 		`gorm:"unique;not null"`
	Password  string
	Token     string
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}