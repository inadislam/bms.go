package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func NotImplemented(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"app name:":    os.Getenv("APP_NAME"),
		"app version:": os.Getenv("APP_VERSION"),
		"app author:":  os.Getenv("APP_AUTHOR"),
		"message:":     "the page you are looking for isn't implemented yet. please revisit.",
	})
}
