package routes

import "github.com/gofiber/fiber/v3"

func Setup(app *fiber.App) {
	SetupFileRoutes(app)
	SetupRegisterRoutes(app)
	SetupLoginRoutes(app)
}
