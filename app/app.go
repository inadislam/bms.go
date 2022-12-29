package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/bms-go/app/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	utils.CheckError(err)
	app := fiber.New()
	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
