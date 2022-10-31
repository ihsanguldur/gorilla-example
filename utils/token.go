package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"time"
	"todo-gorilla/models"
)

func GenerateToken(user *models.User) string {
	claim := models.AccessToken{
		Id:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	token, err := t.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	return token
}
