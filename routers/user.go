package routers

import (
	"github.com/followthepattern/go-simple-webservice/controllers"
	"github.com/followthepattern/go-simple-webservice/services"
	"github.com/gorilla/mux"
)

type UserRouterRegister struct {
	database services.Database
}

func NewUserRouterRegister(database services.Database) *UserRouterRegister {
	return &UserRouterRegister{database: database}
}

func (p *UserRouterRegister) RegisterUserRoutes(router *mux.Router) {
	userController := controllers.NewUser(p.database)
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/users/count", userController.GetUserCount).Methods("GET")
	router.HandleFunc("/users", userController.AddUser).Methods("POST")
	router.HandleFunc("/users", userController.DeleteAllUser).Methods("DELETE")
	router.HandleFunc("/users/{user_id}", userController.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{user_id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{user_id}", userController.DeleteUser).Methods("DELETE")
}
