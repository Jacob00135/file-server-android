package middleware

import (
	"github.com/gofiber/fiber/v3"

	db "github.com/Jacob00135/file-server-android/database"
)

func LoginCheck(c fiber.Ctx) error {
	sess, err := db.Storage.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Could not Get session",
		})
	}

	if name := sess.Get("username"); name == nil {
		return c.Status(fiber.StatusUnauthorized).Render("error", fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "Please login first",
		})
	}
	return c.Next()
}
