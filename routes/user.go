package routes

import (
	"github.com/gorilla/mux"
	"todo-gorilla/controllers"
)

func UserRouter(router *mux.Router) {
	userRouter := router.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("", controllers.CreateUser).Methods("POST")
	userRouter.HandleFunc("/sign-in", controllers.Login).Methods("POST")
}
