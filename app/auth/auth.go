package auth

import (
	"math"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

func IsAuth(c *fiber.Ctx) error {
	token := c.Cookies("access_token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":  "missing or malformed JWT",
			"status": fiber.StatusUnauthorized,
			"data":   nil,
		})
	}
	newToken := strings.Split(token, " ")
	if len(newToken) == 2 {
		tr, err := VerifyToken(newToken[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  "invalid JWT",
				"status": fiber.StatusUnauthorized,
				"data":   nil,
			})
		}
		if claims, ok := tr.Claims.(jwt.MapClaims); ok && tr.Valid {
			sec, dec := math.Modf(claims["exp"].(float64))
			expTime := time.Unix(int64(sec), int64(dec*(1e9)))
			if expTime.Unix() < time.Now().Unix() {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error":  "invalid or expired JWT",
					"status": fiber.StatusUnauthorized,
					"data":   nil,
				})
			}
		}
		return c.Next()
	}
	return nil
}
