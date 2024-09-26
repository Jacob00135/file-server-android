package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	db "github.com/Jacob00135/file-server-android/database"
	"github.com/Jacob00135/file-server-android/models"
)

func RegisterUser(c fiber.Ctx) error {
	user := new(models.UserInput)

	if err := c.Bind().Body(user); err != nil {
		handleRegistrationError(c, err)
	}

	exists, err := db.DB.CheckUserExists(user.Username)
	if err != nil {
		n_err := fmt.Errorf("failed to check if user exists: %w", err)
		return handleRegistrationError(c, n_err)
	}
	if exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	err = db.DB.InsertUser(user.Username, user.Password)
	if err != nil {
		n_err := fmt.Errorf("failed to insert user: %w", err)
		return handleRegistrationError(c, n_err)
	}

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// handleRegistrationError handles errors that occur during user registration.
func handleRegistrationError(c fiber.Ctx, err error) error {
	switch {
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
}
