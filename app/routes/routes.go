package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/inadislam/bms-go/app/auth"
	"github.com/inadislam/bms-go/app/controllers"
)

func NewRoutes(app *fiber.App) {
	app.Use(
		cors.New(cors.Config{
			AllowCredentials: true,
		}),
		compress.New(compress.Config{
			Level: compress.LevelBestCompression,
		}),
	)

	app.Get("/home", auth.IsAuth, controllers.NotImplemented)
	app.Get("/posts", controllers.NotImplemented)
	app.Get("/categories", controllers.NotImplemented)
	app.Get("/category/:catname", controllers.NotImplemented)
	app.Get("/author/:authorname", controllers.NotImplemented)

	app.Post("/register", controllers.Registration)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)
	app.Post("/refresh-token", controllers.NotImplemented)
	app.Post("/active-user/:userid", controllers.ActiveUser)
	app.Post("/search/:q", controllers.NotImplemented)
	app.Post("/comments/:postid", controllers.NotImplemented)
	app.Post("/addpost", controllers.NotImplemented)
	app.Post("/addcategory", controllers.NotImplemented)
	app.Post("/adduser", controllers.NotImplemented)
	app.Post("addcomment", controllers.NotImplemented)
}
