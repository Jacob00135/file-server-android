package routes

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"

	"github.com/Jacob00135/file-server-android/controllers"
)

func SetupLoginRoutes(app *fiber.App) {
	// login
	app.Get("/login", redirectIfLoggedIn)
	app.Post("/api/login", controllers.LoginUser)

	// logout
	app.Get("/logout", controllers.LogoutUser)
}

func redirectIfLoggedIn(c fiber.Ctx) error {
	if IsLoggedIn(c) {
		slog.Info("Redirect to /")
		return c.Redirect().To("/")
	}
	return c.SendFile("frontend/html/login.html")
}

func IsLoggedIn(c fiber.Ctx) bool {
	userp := c.Locals("userPermission").(uint)
	return userp != 1
}
