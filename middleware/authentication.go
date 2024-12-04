package middleware

import (
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

	// get the original product to compare and check if data exist or not
	result := DB.Raw("select * from \"Products\" where id = ?", updatedProduct.ID).Scan(&originalProduct)
	if result.Error != nil {
		return handler.ErrorHandler("internal server error", c)
	} else if result.RowsAffected == 0{
		return handler.ErrorHandler("NOT_FOUND", c)
	}

	//compare original product userId with the login user id
	if originalProduct.UserId != userId{
		return handler.ErrorHandler("UNAUTHORIZED", c)
	}
	return c.Next()
}
