package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string         `gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	if b.ID != "" {
		// If the ID is already set, don't generate a new one
		// This is useful when we need to set the ID ourselves
		return nil
	}

	id, err := uuid.NewV7()
	if err != nil {
		return errors.New("failed to generate UUID: " + err.Error())
	}

	b.ID = id.String()

	return nil
}
