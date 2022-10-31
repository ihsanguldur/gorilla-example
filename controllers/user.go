package controllers

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"todo-gorilla/database"
	"todo-gorilla/models"
	"todo-gorilla/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var err error
	user := new(models.User)

	if err = utils.DecodeBody(r.Body, user); err != nil {
		utils.Error(w, http.StatusBadRequest, "error while parsing body.")
		return
	}

	if msg, ok := utils.ValidateUser(user); !ok {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	if err = database.DB.Create(user).Error; err != nil {
		utils.SysError(w)
		return
	}

	utils.Success(w, nil)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	var err error
	user := new(models.User)

	if err = utils.DecodeBody(r.Body, user); err != nil {
		utils.Error(w, http.StatusBadRequest, "error while parsing body.")
		return
	}

	if msg, ok := utils.ValidateUser(user); !ok {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	if err = database.DB.Where(&models.User{Username: user.Username, Password: user.Password}).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Error(w, http.StatusBadRequest, "username or password is not correct.")
			return
		}
		utils.SysError(w)
		return
	}

	token := utils.GenerateToken(user)

	utils.Success(w, token)
	return
}
