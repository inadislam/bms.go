package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/inadislam/bms-go/app/controllers"
)

func NewRoutes(app *fiber.App) {
	app.Use(cache.New(
		cache.Config{
			Expiration:   24 * time.Hour,
			CacheControl: true,
			CacheHeader:  "X-Cache-Status",
		},
	), cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Get("/home", controllers.NotImplemented)
	app.Get("/posts", controllers.NotImplemented)
	app.Get("/categories", controllers.NotImplemented)
	app.Get("/category/:catname", controllers.NotImplemented)
	app.Get("/author/:authorname", controllers.NotImplemented)

	app.Post("/register", controllers.Registration)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)
	app.Post("/active-user/:userid", controllers.ActiveUser)
	app.Post("/search/:q", controllers.NotImplemented)
	app.Post("/comments/:postid", controllers.NotImplemented)
	app.Post("/addpost", controllers.NotImplemented)
	app.Post("/addcategory", controllers.NotImplemented)
	app.Post("/adduser", controllers.NotImplemented)
	app.Post("addcomment", controllers.NotImplemented)
}
