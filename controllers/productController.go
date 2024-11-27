package controllers

import "github.com/gofiber/fiber/v2"

func GetAllProducts(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": "this is another test",
	})
}
