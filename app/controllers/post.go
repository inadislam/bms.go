package controllers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/inadislam/bms-go/app/auth"
	"github.com/inadislam/bms-go/app/db"
	"github.com/inadislam/bms-go/app/models"
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
	if posts[0].Title == "" {
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

func AddPost(c *fiber.Ctx) error {
	token := c.Cookies("access_token")
	newToken := strings.Split(token, " ")
	post := new(models.Posts)
	var postDetails models.Posts
	if len(newToken) == 2 {
		claims, err := auth.ExtractToken(newToken[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  "unauthorized",
				"status": fiber.StatusUnauthorized,
				"data":   nil,
			})
		}
		if err := c.BodyParser(post); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
				"data":   nil,
			})
		}
		if post.Title == "" || post.Body == "" || post.Status == "" || post.Category == "" || post.FeatureImage == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "field must not be empty!!",
				"status": fiber.StatusBadRequest,
				"data":   nil,
			})
		}
		userid, err := uuid.Parse(fmt.Sprintf("%v", claims["user_id"]))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
				"data":   nil,
			})
		}

		post.AuthorID = userid
		postDetails, err = db.CreatePost(*post, userid.String())
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
		"data":    postDetails,
	})
}

func DeletePost(c *fiber.Ctx) error {
	token := c.Cookies("access_token")
	newToken := strings.Split(token, " ")
	var delete int64
	if len(newToken) == 2 {
		claims, err := auth.ExtractToken(newToken[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  "unauthorized",
				"status": fiber.StatusUnauthorized,
				"data":   nil,
			})
		}
		postId := c.Params("postid")
		delete, err = db.PostDelete(postId, fmt.Sprintf("%v", claims["user_id"]))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "there is error somewhere.",
				"status": fiber.StatusBadRequest,
				"data":   nil,
			})
		}
	}
	if delete > 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "success",
			"data":    delete,
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":  "there is a error somewhere",
		"status": fiber.StatusBadRequest,
		"data":   nil,
	})
}
