package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"

	"github.com/Jacob00135/file-server-android/database"
	"github.com/Jacob00135/file-server-android/routes"
)

func main() {
	// Connect to the database
	database.InitDB()
	defer database.DB.Close()

	// Initialize a new Fiber app
	engin := html.New("./frontend/html", ".html")
	app := fiber.New(fiber.Config{
		Views: engin,
	})

	app.Get("/static/*", static.New("./frontend/static"))

	// app.Get("/frontend/js", static.New("frontend/js"))

	app.Get("/", func(c fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Welcome to Home Page",
		})
	})

	// Define a route for the GET method on the root path '/'
	// app.Get("/", func(c fiber.Ctx) error {
	// 	// Send a string response to the client
	// 	return c.SendString("Welcome file server👋!")
	// })

	routes.SetupFileRoutes(app)
	routes.SetupRegisterRoutes(app)
	routes.SetupLoginRoutes(app)

	// Start the server on port 3000
	log.Fatal(app.Listen("0.0.0.0:9527"))
}
