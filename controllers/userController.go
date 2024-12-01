package controllers

import (
	"shopingList/handler"
	"shopingList/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB
var message string
func Register(c *fiber.Ctx) error {
	var user model.Users
	if err := c.BodyParser(&user); err != nil {
		message = "All form needs to be filled"
		return handler.ErrorHandler(err, c, message)
	}

	if result := DB.Create(&user); result.Error != nil {
		return handler.ErrorHandler(result.Error, c, "Error at register user")
	}
	return c.JSON(fiber.Map{
		"data": "You've succesfully registered",
	})
}
