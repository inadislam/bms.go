package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/inadislam/bms-go/app/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetCategories() (models.Category, error) {
	var categories models.Category
	err := DB.Debug().Model(&models.Category{}).Find(&categories).Error
	if err != nil {
		return models.Category{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Category{}, errors.New("category not found")
	}
	return categories, nil
}

func CategoryById(categoryid uuid.UUID) (models.Category, error) {
	var category models.Category
	err := DB.Debug().Model(&models.Category{}).Where("ID = ?", categoryid).Find(&category).Error
	if err != nil {
		return models.Category{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Category{}, errors.New("post not found")
	}
	return category, nil
}

func CategoryByTitle(categoryTitle string) (models.Category, error) {
	var category models.Category
	err := DB.Debug().Model(&models.Category{}).Where("Category = ?", categoryTitle).Find(&category).Error
	if err != nil {
		return models.Category{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Category{}, errors.New("post not found")
	}
	return category, nil
}

func CreateCategory(category models.Category) (models.Category, error) {
	err := DB.Debug().Model(&models.Category{}).Create(&category).Error
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func CategoryUpdate(category map[string]interface{}, categoryid string) (map[string]interface{}, error) {
	update := DB.Debug().Model(&models.Users{}).Clauses(clause.Returning{}).Where("id = ?", categoryid).Updates(category)
	if update.Error != nil {
		return map[string]interface{}{}, update.Error
	}
	return category, nil
}

func CategoryDelete(category string) (int64, error) {
	categoryid, err := uuid.Parse(category)
	if err != nil {
		return 0, err
	}
	delete := DB.Debug().Where("id = ?", categoryid).Delete(&models.Category{})
	if delete.Error != nil {
		return 0, delete.Error
	}
	return delete.RowsAffected, nil
}
