package main

import (
	"log"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"

	"github.com/Jacob00135/file-server-android/controllers"
	db "github.com/Jacob00135/file-server-android/database"
	"github.com/Jacob00135/file-server-android/middleware"
	"github.com/Jacob00135/file-server-android/routes"
)

func main() {
	// Connect to the database
	db.InitDB()
	defer db.DB.Close()

	// Initialize a new Fiber app
	engin := html.New("./frontend/html", ".html")
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		Views:       engin,
	})

	app.Get("/static/*", static.New("./frontend/static"))

	// app.Get("/frontend/js", static.New("frontend/js"))

	// app.Get("/", func(c fiber.Ctx) error {
	// 	return c.Render("index", fiber.Map{
	// 		"Title": "Welcome to Home Page",
	// 	})
	// })

	// app.Get("/", func(c fiber.Ctx) error {
	// 	return c.SendFile("frontend/html/index.html")
	// })

	app.Get("/", controllers.WebIndex, middleware.FileAuth)

	// Define a route for the GET method on the root path '/'
	// app.Get("/", func(c fiber.Ctx) error {
	// 	// Send a string response to the client
	// 	return c.SendString("Welcome file server👋!")
	// })

	routes.Setup(app)

	// Start the server on port 3000
	log.Fatal(app.Listen("0.0.0.0:9527"))
}
