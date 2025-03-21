package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "secret"

func GenerateJWT(email, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secret))

}
