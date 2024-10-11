package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"

	db "github.com/Jacob00135/file-server-android/database"
	"github.com/Jacob00135/file-server-android/models"
)

func ListUsers(c fiber.Ctx) error {
	users, err := db.DB.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"users":   users,
	})

}

func AddUser(c fiber.Ctx) error {
	user := new(models.UserInput)
	if err := c.Bind().Body(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Could not parse request body",
		})
	}

	err := db.DB.InsertUser(user.Username, user.Password, uint(2))
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("User %s added", user.Username),
	})
}

func DelUser(c fiber.Ctx) error {
	uidStr := c.Params("id")
	uid, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := db.DB.DeleteUserById(uint(uid)); err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("User %s deleted", uidStr),
	})
}

func UpdateUserPwd(c fiber.Ctx) error {
	psw := &struct {
		Password string `json:"password"`
	}{}
	username := c.Locals("username").(string)

	if err := c.Bind().Body(psw); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Could not parse request body",
		})
	}

	ok, err := db.DB.CheckUserExists(username, psw.Password)
	if err != nil {
		n_err := fmt.Errorf("database error, failed to check if user exists: %w", err)
		// return handleRegistrationError(c, n_err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": n_err.Error(),
		})
	}

	if ok {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "New password is the same as the old one",
		})
	}

	err = db.DB.UpdateUser(username, psw.Password, uint(2))
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	sess, err := db.Storage.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Could not Get session",
		})
	}

	// 删除 username，不存在则忽略
	sess.Delete("username")
	err = sess.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Could not save session",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("User %s password updated", username),
	})
}

func ListDirs(c fiber.Ctx) error {
	dirs, err := db.DB.GetAllDir()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"dirs":    dirs,
	})
}

func AddDir(c fiber.Ctx) error {
	NewDir := &struct {
		Path       string
		Permission uint
	}{}

	if err := c.Bind().Body(NewDir); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Could not parse request body",
		})
	}

	err := db.DB.InsertDir(NewDir.Path, NewDir.Permission)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Dir %s added", NewDir.Path),
	})
}

func DelDir(c fiber.Ctx) error {
	uidStr := c.Params("id")
	uid, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := db.DB.DeleteDirById(uint(uid)); err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Dir %s deleted", uidStr),
	})
}

func UpdateDir(c fiber.Ctx) error {
	newDir := &struct {
		DirID      uint
		Path       string
		Permission uint
	}{}

	if err := c.Bind().Body(newDir); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Could not parse request body",
		})
	}

	_, err := db.DB.CheckFileExists(newDir.Path)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	err = db.DB.UpdateDir(newDir.DirID, newDir.Permission, newDir.Path)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Dir %s updated", newDir.Path),
	})
}
