package controllers

import (
	"shopingList/api/helpers"
	"shopingList/api/model"
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
		return helpers.ErrorHandler(errorType, c)
	}

	//check if the form is filled or not
	if user.Password == "" || user.Email == "" || user.Username == "" {
		errorType = "ALL_FORM_REQUIRED"
		return helpers.ErrorHandler(errorType, c)
	}

	//create the new user and send error if meet any validation constraint
	if result := DB.Create(&user); result.Error != nil {

		if strings.Contains(result.Error.Error(), "23505") {
			errorType = "USERNAME_EMAIL_EXIST"
		}
		return helpers.ErrorHandler(errorType, c)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "You've succesfully registered",
	})
}

func Login(c *fiber.Ctx) error {
	var user model.Users
	var foundUser model.Users

	//check if client send the form data
	if err := c.BodyParser(&user); err != nil {
		errorType = "INVALID_BODY"
		return helpers.ErrorHandler(errorType, c)
	}

	//check if the form is filled or not
	if user.Email == "" && user.Username == "" {
		errorType = "EMAIL_OR_USERNAME_REQ"
		return helpers.ErrorHandler(errorType, c)
	} else if user.Password == "" {
		errorType = "PASSWORD_REQ"
		return helpers.ErrorHandler(errorType, c)
	}
	result := DB.Raw("select * from \"Users\" where username = ? or email = ?", user.Username, user.Email).Scan(&foundUser)

	if result.RowsAffected == 0 {
		errorType = "NOT_FOUND"
		return helpers.ErrorHandler(errorType, c)
	}

	checkPass := helpers.ComparePass(user.Password, foundUser.Password)
	if checkPass == nil {
		token, err := helpers.SignToken(foundUser.Username, c)
		if err != nil {
			return helpers.ErrorHandler("internal server error", c)
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data":   "Bearer " + token,
			"userId": foundUser.ID,
		})
	} else {
		errorType = "INVALID_PASSWORD"
		return helpers.ErrorHandler(errorType, c)
	}
}
