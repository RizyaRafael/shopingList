package handler

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(err error, c *fiber.Ctx, message string) error {

	code := fiber.StatusInternalServerError

	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		code = fiberError.Code
		log.Println("error in psg block")
		return c.Status(code).JSON(fiber.Map{
			"message": message,
		})
	}


	log.Println("internal server error")
	log.Println(err.Error())
	return c.Status(code).JSON(fiber.Map{
		"message": "An unexpected error occured",
	})
}
