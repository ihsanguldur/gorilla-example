package controllers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"todo-gorilla/database"
	"todo-gorilla/models"
	"todo-gorilla/utils"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var err error
	todo := new(models.Todo)

	ctx := r.Context()
	user := ctx.Value("user").(*models.AccessToken)

	if err = utils.DecodeBody(r.Body, todo); err != nil {
		utils.Error(w, http.StatusBadRequest, "error while parsing body.")
		return
	}

	todo.UserID = user.Id

	if msg, ok := utils.ValidateTodo(todo); !ok {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	if err = database.DB.Create(todo).Error; err != nil {
		utils.SysError(w)
		return
	}

	utils.Success(w, nil)
	return
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var err error
	var todos []*models.Todo

	ctx := r.Context()
	user := ctx.Value("user").(*models.AccessToken)

	if err = database.DB.Where("user_id = ?", user.Id).Find(&todos).Error; err != nil {
		utils.SysError(w)
		return
	}
	fmt.Println(user.Id)
	utils.Success(w, todos)
	return
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	var err error
	todoID := mux.Vars(r)["todoID"]
	todo := new(models.Todo)

	ctx := r.Context()
	user := ctx.Value("user").(*models.AccessToken)

	if todoID == "" {
		utils.Error(w, http.StatusBadRequest, "todoID required.")
		return
	}

	if err = database.DB.Where("id = ? and user_id = ?", todoID, user.Id).First(todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Error(w, http.StatusBadRequest, "todo not found.")
			return
		}
		utils.SysError(w)
		return
	}

	utils.Success(w, todo)
	return
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var err error
	todoID := mux.Vars(r)["todoID"]
	todo := new(models.Todo)

	ctx := r.Context()
	user := ctx.Value("user").(*models.AccessToken)

	if err = utils.DecodeBody(r.Body, todo); err != nil {
		utils.Error(w, http.StatusBadRequest, "error while parsing body.")
		return
	}

	update := database.DB.Model(&models.Todo{}).Where("id = ? and user_id = ?", todoID, user.Id).
		Updates(models.Todo{Content: todo.Content, Status: todo.Status})

	if update.RowsAffected == 0 {
		utils.Error(w, http.StatusBadRequest, "todo not found.")
		return
	}

	if err = update.Error; err != nil {
		utils.SysError(w)
		return
	}

	utils.Success(w, nil)
	return
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var err error
	todoID := mux.Vars(r)["todoID"]
	todo := new(models.Todo)

	ctx := r.Context()
	user := ctx.Value("user").(*models.AccessToken)

	deleted := database.DB.Where("id = ? and user_id = ?", todoID, user.Id).Delete(todo)

	if deleted.RowsAffected == 0 {
		utils.Error(w, http.StatusBadRequest, "todo not found.")
		return
	}

	if err = deleted.Error; err != nil {
		utils.SysError(w)
		return
	}

	utils.Success(w, nil)
	return
}
