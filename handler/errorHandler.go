package handler

import (
	"github.com/gofiber/fiber/v2"
)
const(
	FormRequiredError = "ALL_FORM_REQUIRED"
	InvalidBodyError = "INVALID_BODY"
	EmailOrUsernameError = "USERNAME_EMAIL_EXIST"
)

func ErrorHandler(err string, c *fiber.Ctx) error {
var statusCode int
var errorMessage string
	switch err {
	case FormRequiredError:
		statusCode = fiber.StatusBadRequest
		errorMessage = "all field cannot be empty"
	case InvalidBodyError:
		statusCode = fiber.StatusInternalServerError
		errorMessage = "invalid body request"
	case EmailOrUsernameError:
		statusCode = fiber.StatusBadRequest
		errorMessage = "email or username already registered"
	default:
		statusCode = fiber.StatusInternalServerError
		errorMessage = "an error occured"
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"message": errorMessage,
	})

}
