package handler

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func SignToken(username string, c *fiber.Ctx) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	log.Println(jwtSecret)
	claims := jwt.MapClaims{
		"username": username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", ErrorHandler("Internal Server Error", c)
	}
	return signedToken, nil
}

func VerifyToken(access_token string, c *fiber.Ctx) (string, error){
	token, err := jwt.Parse(access_token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", ErrorHandler("Internal server error", c)
	}
	claim := token.Claims.(jwt.MapClaims)

	return claim["username"].(string), nil
}
