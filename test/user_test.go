package test

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-gorilla/controllers"
	"todo-gorilla/database"
)

func setup() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()
}

func Test_CreateUser(t *testing.T) {
	setup()

	body := strings.NewReader(`{
		"username": "test_user",
		"password": "test_pass"
	}`)

	req, _ := http.NewRequest("POST", "/users", body)
	w := httptest.NewRecorder()

	controllers.CreateUser(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Create User didn't return %v", http.StatusOK)
	}
}
