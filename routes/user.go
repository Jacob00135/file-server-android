package routes

import (
	"github.com/Jacob00135/file-server-android/middleware"
	"github.com/gofiber/fiber/v3"
)

func SetupUPRoutes(app *fiber.App) {
	app.Get("/api/authentication", func(c fiber.Ctx) error {
		up := c.Locals("userPermission")
		return c.JSON(fiber.Map{
			"success":    true,
			"permission": up,
		})
	}, middleware.GetUserPermission)
}
