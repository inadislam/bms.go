package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/inadislam/bms-go/app/db"
	"github.com/inadislam/bms-go/app/models"
)

func Registration(c *fiber.Ctx) error {
	body := new(models.Users)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(fiber.Map{})
}

func ActiveUser(c *fiber.Ctx) error {
	type Body struct {
		UserId uuid.UUID `json:"user_id"`
		Otp    string    `json:"otp"`
	}
	b := new(Body)
	if err := c.BodyParser(b); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	user, err := db.UserById(b.UserId)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	if b.Otp != user.Verification {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": "internal server error",
		})
	} else {
		if user.Verified == true {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"errors": "otp expired",
			})
		}
		err = db.UserActive(user.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"errors": "failed to active account, internal server error",
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_id":       user.ID,
		"name":          user.Name,
		"email":         user.Email,
		"role":          user.Role,
		"profile_photo": user.ProfilePhoto,
		"password":      "Your Password",
		"message":       "your account activated.please login now!!",
		"status":        fiber.StatusOK,
	})
}
