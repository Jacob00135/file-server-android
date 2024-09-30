package controllers

import (
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v3"

	db "github.com/Jacob00135/file-server-android/database"
	"github.com/Jacob00135/file-server-android/models"
)

func LoginUser(c fiber.Ctx) error {
	slog.Info("Login handler called")
	// Parse the body into the UserInput struct

	// Get session
	sess, err := db.Storage.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Could not Get session",
		})
	}

	// Judge user if already login
	if sess.Get("username") != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "User already login",
		})
	}

	user := new(models.UserInput)
	if err := c.Bind().Body(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Could not parse request body",
		})
	}

	// Check if the user exists in the database and the password is correct
	exists, err := db.DB.CheckUserExists(user.Username, user.Password)
	if err != nil {
		n_err := fmt.Errorf("database error, failed to check if user exists: %w", err)
		// return handleRegistrationError(c, n_err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": n_err.Error(),
		})
	}
	if !exists {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "User or password is incorrect",
		})
	}

	// Store the user's username in the session
	sess.Set("username", user.Username)
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Could not save session",
		})
	}

	// Return a success message
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Successfully logged in",
	})
}
