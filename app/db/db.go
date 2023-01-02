package db

import (
	"fmt"
	"os"

	"github.com/inadislam/bms-go/app/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitDB() {
	con := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	Database, err := gorm.Open(postgres.Open(con), &gorm.Config{})
	defer func() {
		db, err := Database.DB()
		db.Close()
		utils.CheckError(err)
	}()
	utils.CheckError(err)
	fmt.Println("Database Connection Eastablished!!")
}
