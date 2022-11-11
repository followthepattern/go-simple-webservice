package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/followthepattern/go-simple-webservice/models"
	"github.com/followthepattern/go-simple-webservice/services"
	"github.com/followthepattern/go-simple-webservice/utils"
	"github.com/gorilla/mux"
)

type UserController interface {
	GetUsers(http.ResponseWriter, *http.Request)
	GetUserByID(http.ResponseWriter, *http.Request)
	GetUserCount(http.ResponseWriter, *http.Request)
	AddUser(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
	DeleteAllUser(http.ResponseWriter, *http.Request)
}

func NewUser(database services.Database) UserController {
	return &userController{
		database: database,
	}
}

type userController struct {
	database services.Database
}

func (c *userController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := c.database.GetAllUser()
	utils.Write(w, http.StatusAccepted, users)
}

func (c *userController) GetUserCount(w http.ResponseWriter, r *http.Request) {
	count := c.database.GetUserCount()
	utils.Write(w, http.StatusAccepted, count)
}

func (c *userController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]

	if userID == "" {
		utils.BadRequest(w, nil)
		return
	}

	user, err := c.database.GetUserByID(userID)

	if err != nil {
		utils.Write(w, utils.GetErrorCode(err), err.Error())
		return
	}

	utils.Write(w, http.StatusAccepted, user)
}

func (c *userController) AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		utils.BadRequest(w, user)
		return
	}

	c.database.AddUser(user)
	utils.Created(w, user)
	return
}

func (c *userController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	params := mux.Vars(r)

	if err != nil {
		utils.BadRequest(w, user)
		return
	}

	err = c.database.UpdateUser(params["user_id"], user)

	if err != nil {
		utils.Write(w, utils.GetErrorCode(err), err.Error())
		return
	}
	utils.Created(w, user)
	return
}

func (c *userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]
	err := c.database.DeleteUser(userID)
	if err != nil {
		utils.Write(w, utils.GetErrorCode(err), err.Error())
	}
	return
}

func (c *userController) DeleteAllUser(w http.ResponseWriter, r *http.Request) {
	c.database.DeleteAllUser()
	return
}
