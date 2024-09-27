package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/Jacob00135/file-server-android/controllers"
	"github.com/Jacob00135/file-server-android/middleware"
)

func SetupFileRoutes(app *fiber.App) {
	// 上传文件
	app.Post("/upload", controllers.UploadFile)

	// 下载文件
	app.Get("/download/:filename", controllers.DownloadFile)

	app.Get("/api/index", controllers.ListFiles, middleware.FileAuth)

	// app.Get("/api/", ListFiles)
}
