package controllers

import (
	"fmt"
	"strings"
	"time"

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
	var posts models.Posts
	if len(newToken) == 2 {
		claims, err := auth.ExtractToken(newToken[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  "unauthorized",
				"status": fiber.StatusUnauthorized,
				"data":   nil,
			})
		}
		user, err = db.UserByEmail(fmt.Sprintf("%v", claims["email"]))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "user not found",
				"status": fiber.StatusNotFound,
				"data":   nil,
			})
		}
		if name := strings.ReplaceAll(strings.ToLower(user.Name), " ", ""); c.Params("authorname") != name {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "bad request",
				"status": fiber.StatusBadRequest,
				"data":   nil,
			})
		}
		posts, err = db.PostsByUserId(user.ID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "no posts found",
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
			"posts": posts,
		},
	})
}

func UserUpdate(c *fiber.Ctx) error {
	token := c.Cookies("access_token")
	newToken := strings.Split(token, " ")
	users := new(models.UU)
	if err := c.BodyParser(users); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}
	var user map[string]interface{}
	updates := make(map[string]interface{})
	if users.Name != "" {
		updates["name"] = users.Name
	}
	if users.Email != "" {
		updates["email"] = users.Email
	}
	if users.Password != "" {
		updates["password"] = users.Password
	}
	if users.ProfilePhoto != "" {
		updates["profile_photo"] = users.ProfilePhoto
	}
	updates["updated_at"] = time.Now()
	if len(newToken) == 2 {
		claims, err := auth.ExtractToken(newToken[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  "unauthorized",
				"status": fiber.StatusUnauthorized,
				"data":   nil,
			})
		}
		user, err = db.UpdateUser(updates, fmt.Sprintf("%v", claims["user_id"]))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "user not found",
				"status": fiber.StatusNotFound,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success",
		"data":    user,
	})
}
