package controllers

import (
	"shopingList/handler"
	"shopingList/model"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB
var errorType string

func Register(c *fiber.Ctx) error {
	var user model.Users

	//check if client send the form data
	if err := c.BodyParser(&user); err != nil {
		errorType = "INVALID_BODY"
		return handler.ErrorHandler(errorType, c)
	}

	if user.Password == "" || user.Email == "" || user.Username == ""{
		errorType = "ALL_FORM_REQUIRED"
		return handler.ErrorHandler(errorType, c)
	}

	//create the new user and send error if meet any validation constraint
	if result := DB.Create(&user); result.Error != nil {
		
		if strings.Contains(result.Error.Error(), "23505") {
			errorType = "USERNAME_EMAIL_EXIST"
		}
		return handler.ErrorHandler(errorType, c)
	}
	return c.JSON(fiber.Map{
		"data": "You've succesfully registered",
	})
}
