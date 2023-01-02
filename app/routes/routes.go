package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/inadislam/bms-go/app/controllers"
)

func NewRoutes(app *fiber.App) {
	app.Use(cache.New(
		cache.Config{
			Expiration:   24 * time.Hour,
			CacheControl: true,
			CacheHeader:  "X-Cache-Status",
		},
	))

	app.Get("/home", controllers.NotImplemented)
	app.Get("/posts", controllers.NotImplemented)
	app.Get("/categories", controllers.NotImplemented)
	app.Get("/category/:catname", controllers.NotImplemented)
	app.Get("/author/:authorname", controllers.NotImplemented)

	app.Post("/register", controllers.Registration)
	app.Post("/login", controllers.NotImplemented)
	app.Post("/active-user", controllers.NotImplemented)
	app.Post("/search/:q", controllers.NotImplemented)
	app.Post("/comments/:postid", controllers.NotImplemented)
	app.Post("/addpost", controllers.NotImplemented)
	app.Post("/addcategory", controllers.NotImplemented)
	app.Post("/adduser", controllers.NotImplemented)
	app.Post("addcomment", controllers.NotImplemented)
}
