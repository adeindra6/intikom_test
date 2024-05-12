package routes

import (
	"github.com/adeindra6/intikom_test/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
}
