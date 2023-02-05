package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/inadislam/bms-go/app/models"
	"gorm.io/gorm"
)

func GetPosts() (models.Posts, error) {
	var posts models.Posts
	err := DB.Debug().Model(&models.Posts{}).Find(&posts).Error
	if err != nil {
		return models.Posts{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Posts{}, errors.New("user not found")
	}
	return posts, nil
}

func PostsByUserId(userid uuid.UUID) (models.Posts, error) {
	var posts models.Posts
	err := DB.Debug().Model(&models.Posts{}).Where("ID = ?", userid).Find(&posts).Error
	if err != nil {
		return models.Posts{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Posts{}, errors.New("user not found")
	}
	return posts, nil
}
