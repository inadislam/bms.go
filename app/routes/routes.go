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
	app.Get("/posts", controllers.ShowPosts)
	app.Get("/categories", auth.IsAuth, controllers.NotImplemented)
	app.Get("/category/:catname", controllers.NotImplemented)
	app.Get("/author/:authorname", controllers.Author)

	app.Post("/register", controllers.Registration)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)
	app.Post("/refresh-token", controllers.RefreshToken)
	app.Post("/active-user/:userid", controllers.ActiveUser)
	app.Post("/search/:q", controllers.NotImplemented)
	app.Post("/comments/:postid", controllers.NotImplemented)
	app.Post("/addpost", auth.IsAuth, controllers.AddPost)
	app.Post("/addcategory", auth.IsAuth, controllers.NotImplemented)
	app.Post("/adduser", auth.IsAuth, controllers.NotImplemented)
	app.Post("/addcomment", auth.IsAuth, controllers.NotImplemented)
	app.Post("/profile", auth.IsAuth, controllers.UserProfile)
	app.Post("/update-profile", auth.IsAuth, controllers.UserUpdate)
}
