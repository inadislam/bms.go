package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/bms-go/app/db"
	"github.com/inadislam/bms-go/app/routes"
	"github.com/inadislam/bms-go/app/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	utils.CheckError(err)
	db.InitDB()
	db.AutoMigrator()
	app := fiber.New()
	routes.NewRoutes(app)
	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
