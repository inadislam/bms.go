package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/inadislam/bms-go/app/utils"
)

var db *sql.DB

func InitDB() {
	var err error
	con := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	db, err = sql.Open("postgres", con)
	utils.CheckError(err)
	defer db.Close()
	err = db.Ping()
	utils.CheckError(err)
	fmt.Println("Database Connection Eastablished!!")
}
