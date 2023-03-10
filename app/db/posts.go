package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/inadislam/bms-go/app/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetPosts() ([]models.Posts, error) {
	var posts []models.Posts
	err := DB.Debug().Model(&models.Posts{}).Find(&posts).Error
	if err != nil {
		return []models.Posts{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.Posts{}, errors.New("psot not found")
	}
	return posts, nil
}

func GetPostByCatName(catname string) ([]models.Posts, error) {
	var posts []models.Posts
	err := DB.Debug().Model(&models.Posts{}).Where("category = ?", catname).Find(&posts).Error
	if err != nil {
		return []models.Posts{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.Posts{}, errors.New("psot not found")
	}
	return posts, nil
}

func PostsByUserId(userid uuid.UUID) ([]models.Posts, error) {
	var posts []models.Posts
	err := DB.Debug().Model(&models.Posts{}).Where("author_id = ?", userid).Find(&posts).Error
	if err != nil {
		return []models.Posts{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.Posts{}, errors.New("user not found")
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

func UpdatePosts(post map[string]interface{}, postid string) (map[string]interface{}, error) {
	update := DB.Debug().Model(&models.Posts{}).Clauses(clause.Returning{}).Where("id = ?", postid).Updates(post)
	if update.Error != nil {
		return map[string]interface{}{}, update.Error
	}
	return post, nil
}

func PostDelete(postid, userid string) (int64, error) {
	postId, err := uuid.Parse(postid)
	if err != nil {
		return 0, err
	}
	userId, err := uuid.Parse(userid)
	if err != nil {
		return 0, err
	}
	delete := DB.Debug().Model(&models.Posts{}).Where("id = ? AND author_id", postId, userId).Update("visibility", "false")
	if delete.Error != nil {
		return 0, delete.Error
	}
	return delete.RowsAffected, nil
}
