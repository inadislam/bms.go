package db

import (
	"github.com/google/uuid"
	"github.com/inadislam/bms-go/app/models"
	"github.com/inadislam/bms-go/app/utils"
)

func RegistrationHelper(user models.Users) (models.Users, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.Users{}, err
	}
	user.Password = string(hashedPassword)
	err = DB.Debug().Model(&models.Users{}).Create(&user).Error
	if err != nil {
		return models.Users{}, err
	}
	return user, nil
}

func UserById(userid uuid.UUID) (models.Users, error) {
	var user models.Users
	err := DB.Debug().Model(&models.Users{}).Where("ID = ?", userid).Select("id, name, email, password, role, profile_photo, verified").Error
	if err != nil {
		return models.Users{}, err
	}
	return user, nil
}

func UserActive(uid uuid.UUID) error {
	err := DB.Debug().Model(&models.Users{}).Where("ID = ?", uid).Update("verified", true).Error
	if err != nil {
		return err
	}
	return nil
}
