package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Users struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;unique" json:"id"`
	Name         string    `gorm:"type:varchar(255);size:10;not null" json:"name"`
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

type UU struct {
	Name         string    `gorm:"type:varchar(255);size:10;not null" json:"name"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	Password     string    `gorm:"size:500;not null" json:"password,omitempty"`
	ProfilePhoto string    `gorm:"type:varchar(255);not null" json:"profile_photo"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Posts struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key;unique" json:"id"`
	AuthorID     uuid.UUID      `gorm:"type:uuid;not null" json:"author_id"`
	Title        string         `gorm:"type:varchar(255);not null" json:"title"`
	Body         string         `gorm:"not null" json:"body"`
	Status       string         `gorm:"type:varchar(255);not null" json:"status"`
	Category     string         `gorm:"type:varchar(255);not null" json:"category"`
	Tags         string         `gorm:"type:varchar(255);" json:"tags"`
	Medias       pq.StringArray `gorm:"type:text[];" json:"medias"`
	FeatureImage string         `gorm:"type:varchar(255);not null" json:"feature_image"`
	Visibility   bool           `gorm:"default:1;not null" json:"visibility"`
	Publish      time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"publish"`
	CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;unique" json:"id"`
	Category  string    `gorm:"type:varchar(255);unique;not null" json:"category"`
	Slug      string    `gorm:"type:varchar(255);unique;not null" json:"slug"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
