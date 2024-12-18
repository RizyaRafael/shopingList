package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashingPass(password string) (string, error){
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashed), err
}

func ComparePass(password string, comparePass string) error {
	return bcrypt.CompareHashAndPassword([]byte(comparePass), []byte(password))
}