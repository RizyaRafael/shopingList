package middleware

import (
	"log"
	"shopingList/handler"
	"shopingList/model"

	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	var originalProduct model.Products
	var updatedProduct model.Products
	userId := c.Locals("userId")

	// check the request body exist or not
	if err := c.BodyParser(&updatedProduct); err != nil || updatedProduct.ID == 0 {
		errorType := "INVALID_BODY"
		return handler.ErrorHandler(errorType, c)
	}

	// get the original product to compare
	if err := DB.Raw("select * from \"Products\" where id = ?", updatedProduct.ID).Scan(&originalProduct); err.Error != nil {
		
		return handler.ErrorHandler("internal server error", c)
	}

	//compare original product userId with the login user id
	if originalProduct.UserId != userId{
		log.Print("caught in the compare userId")
		return handler.ErrorHandler("UNAUTHORIZED", c)
	}
	return c.Next()
}
