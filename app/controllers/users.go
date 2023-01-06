package controllers

import (
	"strconv"

	"github.com/badoux/checkmail"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/inadislam/bms-go/app/auth"
	"github.com/inadislam/bms-go/app/db"
	"github.com/inadislam/bms-go/app/models"
	"github.com/inadislam/bms-go/app/utils"
)

func Registration(c *fiber.Ctx) error {
	user := new(models.Users)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "fields cannot be empty",
			"status": fiber.StatusBadRequest,
		})
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "invalid email format",
			"status": fiber.StatusUnauthorized,
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
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"ID":           uc.ID,
		"Name":         uc.Name,
		"Email":        uc.Email,
		"Password":     "Your Password",
		"Verification": uc.Verified,
		"Message":      "Check your Email Box for Verification Code",
		"status":       fiber.StatusCreated,
	})
}

func ActiveUser(c *fiber.Ctx) error {
	type Body struct {
		Otp int64 `json:"otp"`
	}
	b := new(Body)
	if err := c.BodyParser(b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}

	userid, _ := uuid.Parse(c.Params("userid"))
	if userid.String() == "00000000-0000-0000-0000-000000000000" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": "user not found",
			"status": fiber.StatusNotFound,
		})
	}
	user, err := db.UserById(userid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": err.Error(),
			"status": fiber.StatusNotFound,
		})
	}
	if b.Otp != user.Verification {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "otp not matched",
			"status": fiber.StatusBadRequest,
		})
	} else {
		if user.Verified {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"errors": "otp expired",
				"status": fiber.StatusInternalServerError,
			})
		}
		err = db.UserActive(user.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"errors": "failed to active account, internal server error",
				"status": fiber.StatusInternalServerError,
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_id":       user.ID,
		"name":          user.Name,
		"email":         user.Email,
		"profile_photo": user.ProfilePhoto,
		"password":      "Your Password",
		"message":       "your account activated.please login now!!",
		"status":        fiber.StatusOK,
	})
}

func Login(c *fiber.Ctx) error {
	user := new(models.Login)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}
	u, err := db.UserByEmail(user.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": "user not found",
			"status": fiber.StatusNotFound,
		})
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "please enter a valid email",
			"status": fiber.StatusUnauthorized,
		})
	}
	if user.Email != u.Email {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "email or password not matched",
			"status": fiber.StatusUnauthorized,
		})
	}
	if err := utils.ComparePass(u.Password, user.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "password not matched",
			"status": fiber.StatusBadRequest,
		})
	}
	if !u.Verified {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "verify your account first!!",
			"status": fiber.StatusUnauthorized,
		})
	}
	token, err := auth.GenerateJWT(u.ID.String(), u.Email)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": "token generating failed",
			"status": fiber.StatusUnprocessableEntity,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": token,
	})
}
