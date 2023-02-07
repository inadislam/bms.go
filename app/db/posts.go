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
	err := DB.Debug().Model(&models.Posts{}).Where("AuthorID = ?", userid).Find(&posts).Error
	if err != nil {
		return models.Posts{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Posts{}, errors.New("user not found")
	}
	return posts, nil
}

func PostsById(postid uuid.UUID) (models.Posts, error) {
	var posts models.Posts
	err := DB.Debug().Model(&models.Posts{}).Where("ID = ?", postid).Find(&posts).Error
	if err != nil {
		return models.Posts{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Posts{}, errors.New("post not found")
	}
	return posts, nil
}

func PostsByTitle(postTitle string) (models.Posts, error) {
	var posts models.Posts
	err := DB.Debug().Model(&models.Posts{}).Where("Title = ?", postTitle).Find(&posts).Error
	if err != nil {
		return models.Posts{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Posts{}, errors.New("post not found")
	}
	return posts, nil
}

func CreatePost(post models.Posts, userid string) (models.Posts, error) {
	err := DB.Debug().Model(&models.Posts{}).Create(&post).Error
	if err != nil {
		return models.Posts{}, err
	}
	return post, nil
}
