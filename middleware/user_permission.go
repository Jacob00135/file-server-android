package middleware

import (
	"github.com/gofiber/fiber/v3"

	db "github.com/Jacob00135/file-server-android/database"
)

func GetUserPermission(c fiber.Ctx) error {
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
		c.Locals("username", username)
	}

	var userPermission uint = 1
	if username != "" {
		var err error
		userPermission, err = db.DB.GetUserPermission(username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
				"code":    fiber.StatusInternalServerError,
				"message": err.Error(),
			})
		}
	}
	// c.Set("X-User-Permission", fmt.Sprintf("%d", userPermission))
	c.Locals("userPermission", userPermission)

	return c.Next()
}
