package handlers

import (
	"go_note/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateSession(userID uint, expiration time.Time, db *gorm.DB) (*models.Session, error) {
	session := &models.Session{
		ID:        uuid.New().String(),
		UserID:    userID,
		ExpiresAt: expiration,
	}

	if err := db.Create(session).Error; err != nil {
		return nil, err
	}

	return session, nil
}
