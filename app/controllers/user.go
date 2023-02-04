package controllers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/bms-go/app/auth"
	"github.com/inadislam/bms-go/app/db"
	"github.com/inadislam/bms-go/app/models"
)

func UserProfile(c *fiber.Ctx) error {
	token := c.Cookies("access_token")
	newToken := strings.Split(token, " ")
	var user models.Users
	if len(newToken) == 2 {
		claims, err := auth.ExtractToken(newToken[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  "unauthorized",
				"status": fiber.StatusUnauthorized,
				"data":   nil,
			})
		}
		user, err = db.UserByEmail(claims["email"].(string))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "user not found",
				"status": fiber.StatusNotFound,
				"data":   nil,
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data": fiber.Map{
			"user_id":       user.ID,
			"email":         user.Email,
			"user_name":     user.Name,
			"profile_photo": user.ProfilePhoto,
		},
	})
}

func Author(c *fiber.Ctx) error {
	token := c.Cookies("access_token")
	newToken := strings.Split(token, " ")
	var user models.Users
	if len(newToken) == 2 {
		claims, err := auth.ExtractToken(newToken[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  "unauthorized",
				"status": fiber.StatusUnauthorized,
				"data":   nil,
			})
		}
		user, err = db.UserByEmail(claims["email"].(string))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "user not found",
				"status": fiber.StatusNotFound,
				"data":   nil,
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data": fiber.Map{
			"user": fiber.Map{
				"user_id":       user.ID,
				"email":         user.Email,
				"user_name":     user.Name,
				"profile_photo": user.ProfilePhoto,
			},
			"posts": fiber.Map{},
		},
	})
}
