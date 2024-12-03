package controllers

import (
	"log"
	"shopingList/handler"
	"shopingList/model"

	"github.com/gofiber/fiber/v2"
)
type Pagination struct {
	Page  int
	Limit int
}

func GetAllProducts(c *fiber.Ctx) error {
	var pagination Pagination
	var response []model.Products
	var totalData int
	if err := c.BodyParser(&pagination); err != nil {
		errorType = "INVALID_BODY"
		return handler.ErrorHandler(errorType, c)
	}
	if pagination.Limit == 0 {
		pagination.Limit = 10
	}

	pagination.Page = pagination.Limit * (pagination.Page - 1)

	DB.Raw("select \"name\", \"id\", \"price\"  from \"Products\" p limit ? offset ?", pagination.Limit, pagination.Page).Scan(&response)
	DB.Raw("select count(*) as total from \"Products\"").Scan(&totalData)

	return c.JSON(fiber.Map{
		"data":  response,
		"total": totalData,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	userId := c.Locals("userId")
	log.Print(userId)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"data": "test",
	})
}


