package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/Jacob00135/file-server-android/controllers"
	"github.com/Jacob00135/file-server-android/middleware"
)

func SetupAdminRoutes(app *fiber.App) {
	app.Get("/api/manage_user", controllers.ListUsers, middleware.AdminAuth)
	// app.Get("/api/administration/:username", controllers.GetUser, middleware.AdminAuth)
	app.Post("/api/manage_user", controllers.AddUser, middleware.AdminAuth)
	app.Put("/api/manage_user/:username", controllers.UpdateUser, middleware.AdminAuth)
	app.Delete("/api/manage_user/:username", controllers.DelUser, middleware.AdminAuth)

	app.Get("/manage_user", func(c fiber.Ctx) error {
		return c.SendFile("frontend/html/manage_user.html")
	}, middleware.AdminAuth)
}
