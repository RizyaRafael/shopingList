package controllers

import (
	"log"
	"shopingList/handler"
	"shopingList/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)



func GetAllProducts(c *fiber.Ctx) error {
	var response []model.Products
	var totalData int

	queries := c.Queries()

	limitStr := queries["limit"]
	pageStr := queries["page"]
	name := queries["name"]

	limit := 10
	page := 1
	var err error
	// check limit value
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			limit = 10
		}
	}

	// check page value
	if pageStr != "" {
		page, err = strconv.Atoi(limitStr)
		if err != nil {
			page = 1
		}
	}

	// check name value
	if name == "" {
		name ="%"
	} else {
		name = "%" + name + "%"
	}

	offset := limit * (page - 1)
	log.Println(name)
	DB.Raw("select \"name\", \"id\", \"price\"  from \"Products\" p where name ilike ? limit ? offset ?",name, limit, offset).Scan(&response)
	DB.Raw("select count(*) as total from \"Products\"").Scan(&totalData)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  response,
		"page":  page,
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
	if newProduct.Name == "" || newProduct.Price == 0 || newProduct.Quantity == 0 || newProduct.ImageUrl == "" {
		errorType = "ALL_FORM_REQUIRED"
		return handler.ErrorHandler(errorType, c)
	}
	newProduct.UserId = userId.(uint)
	result := DB.Create(&newProduct)
	if result.Error != nil {
		return handler.ErrorHandler("DATABASE_ERROR", c)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": newProduct,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	var updatedProduct model.Products
	var originalProduct model.Products

	if err := DB.Raw("Select * from \"Products\" where id = ?", updatedProduct.ID).Scan(&originalProduct); err.Error != nil {
		errorType = "internal server error"
		return handler.ErrorHandler(errorType, c)
	}

	if err := DB.Exec("update \"Products\" set \"name\" = ?, \"price\" = ?, \"quantity\" = ?, \"image_url\" = ? where id = ?", updatedProduct.Name, updatedProduct.Price, updatedProduct.Quantity, updatedProduct.ImageUrl, updatedProduct.ID); err.Error != nil {
		return handler.ErrorHandler("Internal server error", c)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": "Product succesfully updated",
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	var product model.Products

	if err := c.BodyParser(&product); err != nil {
		errorType = "INVALID_BODY"
		return handler.ErrorHandler(errorType, c)
	}

	if err := DB.Exec("delete from \"Products\" where id = ?", product.ID); err.Error != nil {
		return handler.ErrorHandler("internal server error", c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "data succesfully deleted",
	})
}
