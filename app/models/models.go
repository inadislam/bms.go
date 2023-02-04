package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Users struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;unique" json:"id"`
	Name         string    `gorm:"type:varchar(255);not null" json:"name"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	Password     string    `gorm:"size:500;not null" json:"password,omitempty"`
	ProfilePhoto string    `gorm:"type:varchar(255);not null" json:"profile_photo"`
	Verification int64     `gorm:"size:20;default:0;not null" json:"verification,omitempty"`
	Verified     bool      `gorm:"not null" json:"verified,omitempty"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Login struct {
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type Posts struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key;unique" json:"id"`
	AuthorID     uuid.UUID      `gorm:"type:uuid;not null" json:"author_id"`
	Author       Users          `gorm:"foreignKey:AuthorID;" json:"author"`
	Title        string         `gorm:"type:varchar(255);not null" json:"title"`
	Body         string         `gorm:"not null" json:"body"`
	Status       string         `gorm:"type:varchar(255);not null" json:"status"`
	Category     string         `gorm:"type:varchar(255);not null" json:"category"`
	Tags         string         `gorm:"type:varchar(255);not null" json:"tags"`
	Medias       pq.StringArray `gorm:"type:text[];not null" json:"medias"`
	FeatureImage string         `gorm:"type:varchar(255);not null" json:"feature_image"`
	Visibility   bool           `gorm:"type:varchar(255);default:1;not null" json:"visibility"`
	Publish      time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"publish"`
	CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
