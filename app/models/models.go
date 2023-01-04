package models

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;" json:"id"`
	Name         string    `gorm:"type:varchar(255);not null" json:"name"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	Password     string    `gorm:"not null" json:"password"`
	Role         string    `gorm:"type:varchar(255);not null" json:"role"`
	ProfilePhoto string    `gorm:"type:varchar(255);not null" json:"profile_photo"`
	Verification int64     `gorm:"size:20;default:0;not null" json:"verification"`
	Verified     bool      `gorm:"not null" json:"verified"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Login struct {
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
