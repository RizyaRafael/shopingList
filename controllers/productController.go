package controllers

import (
	"shopingList/handler"
	"shopingList/model"

	"github.com/gofiber/fiber/v2"
)

/*
Admin role
1. create product
2. update product
3. delete product

User role
1. Create wishlist
2. Update wishlist
3. Delete wishlist
*/
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
