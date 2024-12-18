package middleware

import (
	"shopingList/handler"
	"shopingList/model"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Authorization(c *fiber.Ctx) error {
	var user model.Users
	access_token := c.Get("Authorization")

	if access_token == "" {
		return handler.ErrorHandler("UNAUTHORIZED", c)
	}

	checkBearer := strings.Split(access_token, " ")
	if checkBearer[0] != "Bearer" {
		return handler.ErrorHandler("UNAUTHORIZED", c)
	}

	verifyToken, err := handler.VerifyToken(checkBearer[1], c)
	if err != nil {
		return handler.ErrorHandler("internal server error", c)
	}
	checkUsername := DB.Raw("select * from \"Users\" where \"username\" = ?", verifyToken).Scan(&user)
	if checkUsername.RowsAffected == 0 {
		return handler.ErrorHandler("UNAUTHORIZED", c)
	}

	c.Locals("userId", user.ID)

	return c.Next()
}
