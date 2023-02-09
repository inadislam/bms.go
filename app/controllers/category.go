package controllers

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/bms-go/app/db"
	"github.com/inadislam/bms-go/app/models"
)

func ShowCategories(c *fiber.Ctx) error {
	cat, err := db.GetCategories()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  "no posts found",
			"status": fiber.StatusNotFound,
			"data":   nil,
		})
	}
	if cat.Category == "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": fiber.StatusOK,
			"data":   "no category found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   cat,
	})
}

func AddCategory(c *fiber.Ctx) error {
	token := c.Cookies("access_token")
	newToken := strings.Split(token, " ")
	category := new(models.Category)
	var categoryDetails models.Category
	if len(newToken) == 2 {
		if err := c.BodyParser(category); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
				"data":   nil,
			})
		}
		if category.Category == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "field must not be empty!!",
				"status": fiber.StatusBadRequest,
				"data":   nil,
			})
		}
		if category.Slug == "" {
			if strings.Contains(category.Category, " ") {
				category.Slug = strings.ToLower(strings.Join(strings.Split(category.Category, " "), "-"))
			} else {
				category.Slug = strings.ToLower(category.Category)
			}
		}

		var err error
		categoryDetails, err = db.CreateCategory(*category)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
				"data":   nil,
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success",
		"data":    categoryDetails,
	})
}

func CategoryByName(c *fiber.Ctx) error {
	catname := c.Params("catname")
	category, err := db.CategoryByTitle(catname)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusNotFound,
			"data":   nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success",
		"data":    category,
	})
}

func UpdateCategory(c *fiber.Ctx) error {
	token := c.Cookies("access_token")
	newToken := strings.Split(token, " ")
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}
	var cat map[string]interface{}
	var err error
	updates := make(map[string]interface{})
	if category.Category != "" {
		updates["category"] = category.Category
	}
	if category.Slug != "" {
		updates["slug"] = category.Slug
	}
	updates["updated_at"] = time.Now()
	if len(newToken) == 2 {
		cat, err = db.CategoryUpdate(updates, c.Params("catid"))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "category not found",
				"status": fiber.StatusNotFound,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success",
		"data":    cat,
	})
}
