package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"github.com/Jacob00135/file-server-android/database"
	"github.com/Jacob00135/file-server-android/models"
)

func SetupLoginRoutes(app *fiber.App) {
	app.Post("/login", loginUser)
}

func loginUser(c fiber.Ctx) error {
	// Parse the body into the UserInput struct
	user := new(models.UserInput)
	if err := c.Bind().Body(user); err != nil {
		handleRegistrationError(c, err)
	}

	// Check if the user exists in the database and the password is correct
	exists, err := database.DB.CheckUserExists(user.Username)
	if err != nil {
		n_err := fmt.Errorf("failed to check if user exists: %w", err)
		return handleRegistrationError(c, n_err)
	}
	if exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	// Get session
	sess, err := database.Storage.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not Get session",
		})
	}

	// Store the user's username in the session
	sess.Set("username", user.Username)
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not save session",
		})
	}

	// Return a success message
	return c.JSON(fiber.Map{
		"message": "Successfully logged in",
	})
}
