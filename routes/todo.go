package routes

import (
	"github.com/gorilla/mux"
	"todo-gorilla/controllers"
	"todo-gorilla/middleware"
)

func TodoRouter(router *mux.Router) {
	todoRouter := router.PathPrefix("/todos").Subrouter()

	todoRouter.Use(middleware.Protected)

	todoRouter.HandleFunc("", controllers.CreateTodo).Methods("POST")
	todoRouter.HandleFunc("", controllers.GetTodos).Methods("GET")
	todoRouter.HandleFunc("/{todoID}", controllers.GetTodo).Methods("GET")
	todoRouter.HandleFunc("/{todoID}", controllers.UpdateTodo).Methods("PUT")
	todoRouter.HandleFunc("/{todoID}", controllers.DeleteTodo).Methods("DELETE")

}
