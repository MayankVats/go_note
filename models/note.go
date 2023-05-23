package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UserID    uint
	Title     string
	Content   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
