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
		log.Print(err)
		return "INTERNAL_ERROR"
	}
	return signedToken
}