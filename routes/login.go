package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/Jacob00135/file-server-android/controllers"
)

func SetupLoginRoutes(app *fiber.App) {
	app.Get("/login", func(c fiber.Ctx) error {
		return c.SendFile("frontend/html/login.html")
	})

	app.Post("/api/login", controllers.LoginUser)
}
