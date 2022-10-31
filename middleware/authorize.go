package middleware

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"todo-gorilla/models"
	"todo-gorilla/utils"
)

func Protected(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			utils.Error(w, http.StatusUnauthorized, "Sign-in please.")
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &models.AccessToken{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if token.Valid {
			ctx := context.WithValue(r.Context(), "user", token.Claims)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			utils.Error(w, http.StatusNotAcceptable, "that's not even a token")
			return
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			utils.Error(w, http.StatusUnauthorized, "timing is everything")
			return
		} else {
			utils.Error(w, http.StatusUnauthorized, err.Error())
			return
		}
	})
}
