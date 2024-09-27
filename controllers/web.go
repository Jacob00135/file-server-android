package controllers

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v3"
)

func WebIndex(c fiber.Ctx) error {

	target := c.Locals("target")
	if target == nil {
		slog.Info("SendFile: frontend/html/index.html")
		return c.SendFile("frontend/html/index.html")
	}

	path := target.(string)
	targeInfo, _ := os.Stat(path)
	if !targeInfo.IsDir() {
		slog.Info("SendFile: " + path)
		return c.Download(path)
	}

	slog.Info("SendFile: frontend/html/index.html")
	return c.SendFile("frontend/html/index.html")
}
