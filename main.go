package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"todo-gorilla/database"
	"todo-gorilla/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	router := mux.NewRouter()

	routes.UserRouter(router)
	routes.TodoRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
