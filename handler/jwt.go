package handler

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func SignToken(username string) string {
	jwtSecret := os.Getenv("JWT_SECRET")
	log.Println(jwtSecret)
	claims := jwt.MapClaims{
		"username": username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "INTERNAL_ERROR"
	}
	return signedToken
}

func VerifyToken(access_token string) string{
	token, err := jwt.Parse(access_token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "INTERNAL_ERROR"
	}
	claim := token.Claims.(jwt.MapClaims)

	return claim["username"].(string)
}
