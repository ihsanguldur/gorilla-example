package utils

import (
	"todo-gorilla/models"
)

func ValidateUser(user *models.User) (string, bool) {
	if user.Username == "" {
		return "username required.", false
	} else if user.Password == "" {
		return "password required.", false
	} else if user.Password != "" && len(user.Password) < 6 {
		return "password must be 6 characters.", false
	}
	return "", true
}

func ValidateTodo(todo *models.Todo) (string, bool) {
	if todo.Content == "" {
		return "content required.", false
	}
	return "", true
}
