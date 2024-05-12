package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/adeindra6/intikom_test/pkg/models"
	"github.com/adeindra6/intikom_test/pkg/utils"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var NewUser models.User

type ErrMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	c := CreateUser.CreateUser()
	res, err := json.Marshal(c)
	if err != nil {
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error while creating new user",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	newUser := models.GetAllUsers()
	res, err := json.Marshal(newUser)
	if err != nil {
		fmt.Print("Error when fetching all users")
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error when fetching all users",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	UserDetails, _ := models.GetUserById(id)
	res, err := json.Marshal(UserDetails)
	if err != nil {
		fmt.Println("Error when fetching user")
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error when fetching user",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]

	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error when updating user")
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error when updating user",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	userDetails, db := models.GetUserById(id)
	if updateUser.Name != "" {
		userDetails.Name = updateUser.Name
	}
	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email
	}
	if updateUser.Password != "" {
		newPassword, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		userDetails.Password = string(newPassword)
	}

	db.Save(&userDetails)
	res, err := json.Marshal(userDetails)
	if err != nil {
		fmt.Println("Error while parsing!!!")
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error when deleting user")
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error when deleting user",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	user := models.DeleteUserById(id)
	res, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error while parsing!!!")
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	Login := &models.User{}
	utils.ParseBody(r, Login)
	c := Login.Login()
	res, err := json.Marshal(c)
	if err != nil {
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Wrong Email or Password!",
			Code:    http.StatusUnauthorized,
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err_msg)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
