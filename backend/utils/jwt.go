package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET = []byte("SECRET_KEY")

func GenerateToken(userID int, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET)
}