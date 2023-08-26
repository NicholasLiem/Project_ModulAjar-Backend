package user

import (
	"encoding/json"
	response "github.com/NicholasLiem/GoLang_Microservice/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CreateUserDTO struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func HelloHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(rw).Encode("Hello")
	if err != nil {
		return
	}
}

func CreateUserHandler(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.ErrorResponse(rw, http.StatusBadRequest, "Invalid user ID: "+err.Error())
		return
	}

	var newUser CreateUserDTO
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Fail to decode user: "+err.Error())
		return
	}

	userModel := UserModel{
		UserID:   uint(userID),
		Username: newUser.Username,
		Email:    newUser.Email,
	}

	err = userModel.setPassword(newUser.Password) // Set password and hash it
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Failed to set password: "+err.Error())
		return
	}
	err = CreateUser(&userModel)
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Failed to create user: "+err.Error())
		return
	}

	response.SuccessResponse(rw, http.StatusCreated, "User created successfully", userModel)
	return
}
