package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID        string    `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
