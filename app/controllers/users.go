package controllers

import (
	"strconv"

	"github.com/badoux/checkmail"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/inadislam/bms-go/app/db"
	"github.com/inadislam/bms-go/app/models"
	"github.com/inadislam/bms-go/app/utils"
)

func Registration(c *fiber.Ctx) error {
	user := new(models.Users)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": "fields cannot be empty",
			"status": fiber.StatusUnprocessableEntity,
		})
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": "invalid email format",
			"status": fiber.StatusUnprocessableEntity,
		})
	}
	uc, err := db.RegistrationHelper(*user)
	if err == nil {
		code := db.GetOTP(uc.ID)
		otp := strconv.FormatInt(code, 10)
		utils.ActiveUser(otp, uc.Email, uc.Name)
	} else {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"errors": "username or email already exist",
			"status": fiber.StatusConflict,
		})
	}
	return c.JSON(fiber.Map{
		"ID":           uc.ID,
		"Name":         uc.Name,
		"Email":        uc.Email,
		"Password":     "Your Password",
		"Verification": uc.Verified,
		"Message":      "Check your Email Box for Verification Code",
		"status":       fiber.StatusOK,
	})
}

func ActiveUser(c *fiber.Ctx) error {
	type Body struct {
		UserId uuid.UUID `json:"user_id"`
		Otp    int64     `json:"otp"`
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
		// err = db.UserActive(user.ID)
		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 		"errors": "failed to active account, internal server error",
		// 	})
		// }
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
