package controllers

import (
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
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
	}

	// check name value
	if name == "" {
		name = "%"
	} else {
		name = "%" + name + "%"
	}

	offset := limit * (page - 1)
	DB.Raw("select \"name\", \"id\", \"price\", \"quantity\", \"user_id\", \"image_url\"  from \"Products\" p where name ilike ? order by name asc limit ? offset ? ", name, limit, offset).Scan(&response)
	DB.Raw("select count(*) as total from \"Products\" where name ilike ?", name).Scan(&totalData)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  response,
		"page":  page,
		"total": totalData,
	})
}

func GetOneProduct(c *fiber.Ctx) error {
	var data model.Products
	productId := c.Params("id")

	result := DB.Raw("select * from \"Products\" where id = ?", productId).Scan(&data)
	if result.Error != nil {
		return handler.ErrorHandler("NOT_FOUND", c)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": data,
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
	result := DB.Exec("insert into \"Products\" (name, quantity, price, image_url, user_id) values (?, ?, ?, ?, ?)", newProduct.Name, newProduct.Quantity, newProduct.Price, newProduct.ImageUrl, userId)
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
	productId := c.Params("id")

	if err := c.BodyParser(&updatedProduct); err != nil {
		errorType = "INVALID_BODY"
		return handler.ErrorHandler(errorType, c)
	}

	if err := DB.Raw("Select * from \"Products\" where id = ?", productId).Scan(&originalProduct); err.Error != nil {
		errorType = "internal server error"
		return handler.ErrorHandler(errorType, c)
	}

	response := DB.Exec("update \"Products\" set \"name\" = ?, \"price\" = ?, \"quantity\" = ?, \"image_url\" = ? where id = ?", updatedProduct.Name, updatedProduct.Price, updatedProduct.Quantity, updatedProduct.ImageUrl, productId)
	if response.Error != nil {
		return handler.ErrorHandler("Internal server error", c)
	}
	if response.RowsAffected == 0 {
		return handler.ErrorHandler("UPDATE_FAILED", c)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": "Product succesfully updated",
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	productId := c.Params("id")
	if err := DB.Exec("delete from \"Products\" where id = ?", productId); err.Error != nil {
		return handler.ErrorHandler("internal server error", c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "data succesfully deleted",
	})
}

func BuyProduct(c *fiber.Ctx) error {
	var bodyProduct model.Products
	var checkProduct model.Products
	if err := c.BodyParser(&bodyProduct); err != nil {
		errorType = "INVALID_BODY"
		return handler.ErrorHandler(errorType, c)
	}
	result := DB.Raw("select * from \"Products\" where id = ?", bodyProduct.ID).Scan(&checkProduct)
	if result.Error != nil {
		return handler.ErrorHandler("internal server error", c)
	} else if result.RowsAffected == 0 {
		return handler.ErrorHandler("NOT_FOUND", c)
	}
	if bodyProduct.Quantity < 1 {
		return handler.ErrorHandler("INVALID_QUANTITY", c)
	}
	update := DB.Exec("update \"Products\" set quantity = ? where id = ?", bodyProduct.Quantity-1, bodyProduct.ID)
	if update.RowsAffected == 0 {
		return handler.ErrorHandler("internal server error", c)
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": "Successfully bought item",
		})
	}
}

func GetUserProducts(c *fiber.Ctx) error {
	var userProduct []model.Products
	userId := c.Locals("userId")
	result := DB.Raw("select * from \"Products\" where user_id = ? order by id asc", userId).Scan(&userProduct)
	if result.Error != nil {
		return handler.ErrorHandler("internal server error", c)
	}

	if len(userProduct) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": "You have not made any product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": userProduct,
	})

}
