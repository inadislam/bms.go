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
}
