package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/Jacob00135/file-server-android/controllers"
)

func SetupRegisterRoutes(app *fiber.App) {
	app.Post("/register", controllers.RegisterUser)
}
