package routes

import (
	"github.com/Jacob00135/file-server-android/controllers"
	"github.com/Jacob00135/file-server-android/middleware"
	"github.com/gofiber/fiber/v3"
)

func SetupUserRoutes(app *fiber.App) {
	app.Get("/api/authentication", func(c fiber.Ctx) error {
		up := c.Locals("userPermission")
		return c.JSON(fiber.Map{
			"success":    true,
			"permission": up,
		})
	}, middleware.GetUserPermission)

	app.Post("/api/change_password", controllers.UpdateUserPwd, middleware.LoginCheck)
	app.Get("/change_password", func(c fiber.Ctx) error {
		return c.SendFile("frontend/html/change_password.html")
	}, middleware.LoginCheck)

	// app.Get("/change_password",)

}
