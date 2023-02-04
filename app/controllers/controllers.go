package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/bms-go/app/auth"
)

func NotImplemented(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"app name:":    os.Getenv("APP_NAME"),
		"app version:": os.Getenv("APP_VERSION"),
		"app author:":  os.Getenv("APP_AUTHOR"),
		"message:":     "the page you are looking for isn't implemented yet. please revisit.",
	})
}

func RefreshToken(c *fiber.Ctx) error {
	RefreshToken := c.Cookies("refresh_token")
	token, err := auth.ExtractTokenAuth(RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusUnauthorized,
		})
	}
	at := fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
	}
	c.Cookie(&at)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"status":  fiber.StatusOK,
	})
}
