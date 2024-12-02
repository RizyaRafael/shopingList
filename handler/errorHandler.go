package handler

import (
	"github.com/gofiber/fiber/v2"
)
const(
	FormRequiredError = "ALL_FORM_REQUIRED"
	InvalidBodyError = "INVALID_BODY"
	EmailOrUsernameError = "USERNAME_EMAIL_EXIST"
	EmptyEmailOrUsername ="EMAIL_OR_USERNAME_REQ"
	PasswordRequiredError = "PASSWORD_REQ"
	DataNotFound = "NOT_FOUND"
	InvalidPasswordError = "INVALID_PASSWORD"
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
	case EmptyEmailOrUsername:
		statusCode = fiber.StatusBadRequest
		errorMessage = "email or username required"
	case PasswordRequiredError:
		statusCode = fiber.StatusBadRequest
		errorMessage = "password is required"
	case DataNotFound:
		statusCode = fiber.StatusNotFound
		errorMessage = "data not found"
	case InvalidPasswordError:
		statusCode = fiber.StatusUnauthorized
		errorMessage = "invalid password or username"
	default:
		statusCode = fiber.StatusInternalServerError
		errorMessage = "an error occured"
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"message": errorMessage,
	})

}
