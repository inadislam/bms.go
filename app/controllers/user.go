package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/inadislam/bms-go/app/auth"
	"github.com/inadislam/bms-go/app/db"
	"github.com/inadislam/bms-go/app/models"
	"github.com/inadislam/bms-go/app/utils"
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
	var user models.Users
	var posts []models.Posts
	userid := c.Params("authorid")
	userId, err := uuid.Parse(userid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "bad request",
			"status": fiber.StatusBadRequest,
			"data":   nil,
		})
	}
	posts, err = db.PostsByUserId(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  "no posts found",
			"status": fiber.StatusNotFound,
			"data":   nil,
		})
	}
	user, err = db.UserById(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  "no posts found",
			"status": fiber.StatusNotFound,
			"data":   nil,
		})
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
	var user map[string]interface{}
	if len(newToken) == 2 {
		claims, err := auth.ExtractToken(newToken[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  "unauthorized",
				"status": fiber.StatusUnauthorized,
				"data":   nil,
			})
		}
		userid := fmt.Sprintf("%v", claims["user_id"])
		if err := c.BodyParser(users); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusBadRequest,
			})
		}
		updates := make(map[string]interface{})
		if users.Name != "" {
			updates["name"] = users.Name
		}
		if users.Email != "" {
			updates["email"] = users.Email
		}
		if users.Password != "" && users.OldPassword != "" {
			id, _ := uuid.Parse(userid)
			u, err := db.UserById(id)
			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error":  "user not found",
					"status": fiber.StatusNotFound,
				})
			}
			err = utils.ComparePass(u.Password, users.OldPassword)
			if err == nil {
				hashedPassword, err := utils.HashPassword(users.Password)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error":  err.Error(),
						"status": fiber.StatusInternalServerError,
					})
				}
				updates["password"] = string(hashedPassword)
			} else {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":  err.Error(),
					"status": fiber.StatusBadRequest,
				})
			}
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "fields must not be empty",
				"status": fiber.StatusBadRequest,
			})
		}
		if users.ProfilePhoto != "" {
			updates["profile_photo"] = users.ProfilePhoto
		}
		updates["updated_at"] = time.Now()
		user, err = db.UpdateUser(updates, userid)
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
