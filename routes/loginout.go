package routes

import (
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v3"

	"github.com/Jacob00135/file-server-android/controllers"
)

func SetupLoginRoutes(app *fiber.App) {
	app.Get("/login", redirectIfLoggedIn)
	app.Post("/api/login", controllers.LoginUser)

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
	fmt.Println(userp)
	return userp != 1
}
