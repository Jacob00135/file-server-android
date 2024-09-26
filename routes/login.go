package routes

import (
	"github.com/Jacob00135/file-server-android/controllers"
	"github.com/gofiber/fiber/v3"
)

func SetupLoginRoutes(app *fiber.App) {
	app.Post("/login", controllers.LoginUser)
}
