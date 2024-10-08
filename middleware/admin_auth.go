package middleware

import (
	"github.com/gofiber/fiber/v3"

	db "github.com/Jacob00135/file-server-android/database"
)

func AdminAuth(c fiber.Ctx) error {
	sess, err := db.Storage.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Could not Get session",
		})
	}

	var username string
	if name := sess.Get("username"); name != nil {
		username = name.(string)
	}

	if username != "admin" {
		return c.Status(fiber.StatusForbidden).Render("error", fiber.Map{
			"code":    fiber.StatusForbidden,
			"message": "Permission denied",
		})
	}

	return c.Next()
}
