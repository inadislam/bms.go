package db

import (
	"fmt"
	"os"

	"github.com/inadislam/bms-go/app/models"
	"github.com/inadislam/bms-go/app/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	con := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	DB, err = gorm.Open(postgres.Open(con), &gorm.Config{})
	utils.CheckError(err)
	fmt.Println("DB Connection Eastablished!!")
}

func AutoMigrator() {
	// err := DB.Debug().Migrator().DropTable(&models.Users{})
	// utils.CheckError(err)
	err := DB.Debug().AutoMigrate(&models.Users{}, &models.Posts{})
	utils.CheckError(err)
}
