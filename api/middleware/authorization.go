package middleware

import (
	"shopingList/api/helpers"
	"shopingList/api/model"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Authorization(c *fiber.Ctx) error {
	var user model.Users
	access_token := c.Get("Authorization")

	if access_token == "" {
		return helpers.ErrorHandler("UNAUTHORIZED", c)
	}

	checkBearer := strings.Split(access_token, " ")
	if checkBearer[0] != "Bearer" {
		return helpers.ErrorHandler("UNAUTHORIZED", c)
	}

	verifyToken, err := helpers.VerifyToken(checkBearer[1], c)
	if err != nil {
		return helpers.ErrorHandler("internal server error", c)
	}
	checkUsername := DB.Raw("select * from \"Users\" where \"username\" = ?", verifyToken).Scan(&user)
	if checkUsername.RowsAffected == 0 {
		return helpers.ErrorHandler("UNAUTHORIZED", c)
	}

	c.Locals("userId", user.ID)

	return c.Next()
}
