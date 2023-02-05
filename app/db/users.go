package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/inadislam/bms-go/app/models"
	"github.com/inadislam/bms-go/app/utils"
	"gorm.io/gorm"
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
	err := DB.Debug().Model(&models.Users{}).Where("ID = ?", userid).Select("id, name, email, password, profile_photo, verification, verified").Find(&user).Error
	if err != nil {
		return models.Users{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Users{}, errors.New("user not found")
	}
	return user, nil
}

func UserByEmail(email string) (models.Users, error) {
	var user models.Users
	err := DB.Debug().Model(&models.Users{}).Where("email = ?", email).Select("id, name, password, email, profile_photo, verification, verified").Find(&user).Error
	if err != nil {
		return models.Users{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Users{}, errors.New("user not found")
	}
	return user, nil
}

func UserActive(uid uuid.UUID) error {
	err := DB.Debug().Model(&models.Users{}).Where("ID = ?", uid).Update("verified", true).Error
	if err != nil {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	}
	return nil
}

func GetOTP(uid uuid.UUID) int64 {
	code, err := utils.GenerateOTP()
	if err != nil {
		return 0
	}
	err = DB.Debug().Model(&models.Users{}).Where("ID = ?", uid).Update("verification", code).Error
	if err != nil {
		return 0
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0
	}
	return code
}

func Users() (models.Users, error) {
	var users models.Users
	err := DB.Debug().Model(&models.Users{}).Find(&users).Error
	if err != nil {
		return models.Users{}, err
	}
	return users, nil
}

func UpdateUser(user models.Users, userid string) (models.Users, error) {
	if user.Password != "" {
		id, _ := uuid.Parse(userid)
		u, err := UserById(id)
		if err != nil {
			return models.Users{}, err
		}
		err = utils.ComparePass(u.Password, user.Password)
		if err == nil {
			hashedPassword, err := utils.HashPassword(user.Password)
			if err != nil {
				return models.Users{}, err
			}
			user.Password = string(hashedPassword)
		} else {
			return models.Users{}, err
		}
	}
	update := DB.Debug().Model(&models.Users{}).Where("id = ?", userid).Updates(
		map[string]interface{}{
			"name":          user.Name,
			"email":         user.Email,
			"password":      user.Password,
			"profile_photo": user.ProfilePhoto,
		},
	)
	if update.Error != nil {
		return models.Users{}, update.Error
	}
	return user, nil
}
