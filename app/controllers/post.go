package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/bms-go/app/db"
)

func ShowPosts(c *fiber.Ctx) error {
	posts, err := db.GetPosts()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  "no posts found",
			"status": fiber.StatusNotFound,
			"data":   nil,
		})
	}
	if posts.Title == "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": fiber.StatusOK,
			"data":   "no post found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   posts,
	})
}
