package middleware

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v3"

	db "github.com/Jacob00135/file-server-android/database"
)

func FileAuth(c fiber.Ctx) error {
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
		fmt.Println("username: ", username, "userPermission: ", userPermission)
	}
	c.Locals("userPermission", userPermission)

	dir := filepath.Clean(c.Query("visible_dir"))

	if dir == "." {
		return c.Next()
	}

	dirPermission, err := db.DB.GetFilePermission(dir)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	if userPermission < dirPermission {
		return c.Status(fiber.StatusForbidden).Render("error", fiber.Map{
			"code":    fiber.StatusForbidden,
			"message": "visible_dir permission denied",
		})
	}

	path := filepath.Clean(c.Query("path"))
	target := filepath.Join(dir, path)
	if !securePath(dir, target) {
		return c.Status(fiber.StatusForbidden).Render("error", fiber.Map{
			"code":    fiber.StatusForbidden,
			"message": "Invalid path",
		})
	}
	if !osFileExists(target) {
		return c.Status(fiber.StatusNotFound).Render("error", fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "File not found",
		})
	}

	c.Locals("target", target)
	return c.Next()
}

func securePath(base, full string) bool {
	absBase, _ := filepath.Abs(base)
	absFull, _ := filepath.Abs(full)
	return strings.HasPrefix(absFull, absBase)
}

func osFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
