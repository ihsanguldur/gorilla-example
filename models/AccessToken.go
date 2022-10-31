package models

import "github.com/golang-jwt/jwt/v4"

type AccessToken struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
