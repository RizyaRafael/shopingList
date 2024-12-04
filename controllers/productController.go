package controllers

import (
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
		"page": pagination.Page,
		"total": totalData,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var newProduct model.Products
	userId := c.Locals("userId")

	if err := c.BodyParser(&newProduct); err != nil {
		errorType = "INVALID_BODY"
		return handler.ErrorHandler(errorType, c)
	}
	if newProduct.Name == "" || newProduct.Price == 0 || newProduct.Quantity == 0 || newProduct.ImageUrl == ""{
		errorType = "ALL_FORM_REQUIRED"
		return handler.ErrorHandler(errorType, c)
	}
	newProduct.UserId = userId.(uint)
	if result := DB.Create(&newProduct); result.Error != nil {
        return handler.ErrorHandler("DATABASE_ERROR", c)
    }

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": "Product succesfully created",
	})
}
