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

func UpdateUser(c fiber.Ctx) error {
	user := struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Permission uint   `json:"permission"`
	}{}

	if err := c.Bind().Body(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Could not parse request body",
		})
	}

	err := db.DB.UpdateUser(user.Username, user.Password, user.Permission)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("User %s updated", user.Username),
	})
}

func ListDirs() {
}

func AddDir(dir string, permission uint) {
}

func DelDir() {
}

func UpdateDir() {
}
