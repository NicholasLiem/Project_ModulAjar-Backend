package user

import (
	"encoding/json"
	"errors"
	response "github.com/NicholasLiem/GoLang_Microservice/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateUserHandler(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	userID, err := verifyUserId(id)
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

	err = userModel.setPassword(newUser.Password)
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

func FindUserByIdHandler(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	userID, err := verifyUserId(id)
	if err != nil {
		response.ErrorResponse(rw, http.StatusBadRequest, "Invalid user ID: "+err.Error())
		return
	}

	condition := UserModel{UserID: uint(userID)}

	foundUser, err := FindOneUser(condition)
	if err != nil {
		response.ErrorResponse(rw, http.StatusNotFound, "User not found")
		return
	}

	response.SuccessResponse(rw, http.StatusOK, "User found", foundUser)
	return
}

func DeleteUserByIdHandler(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	userID, err := verifyUserId(id)
	if err != nil {
		response.ErrorResponse(rw, http.StatusBadRequest, "Invalid user ID: "+err.Error())
		return
	}

	condition := UserModel{UserID: uint(userID)}

	foundUser, err := FindOneUser(condition)
	if err != nil {
		response.ErrorResponse(rw, http.StatusNotFound, "User not found")
		return
	}

	userData := foundUser

	err = foundUser.Delete()
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Fail to delete user")
		return
	}

	response.SuccessResponse(rw, http.StatusOK, "User deleted", userData)
	return
}

func UpdateUserHandler(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	userID, err := verifyUserId(id)
	if err != nil {
		response.ErrorResponse(rw, http.StatusBadRequest, "Invalid user ID: "+err.Error())
		return
	}

	condition := UserModel{UserID: uint(userID)}

	foundUser, err := FindOneUser(condition)
	if err != nil {
		response.ErrorResponse(rw, http.StatusNotFound, "User not found")
		return
	}

	var updateData UpdateUserDTO
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Fail to decode update data: "+err.Error())
		return
	}

	err = foundUser.Update(updateData)
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Fail to update update data: "+err.Error())
		return
	}

	response.SuccessResponse(rw, http.StatusOK, "Updated user data", updateData)
}

func verifyUserId(UserID string) (uint64, error) {
	userID, err := strconv.ParseUint(UserID, 10, 64)
	if err != nil {
		return 0, errors.New("cannot parse id")
	}
	return userID, nil
}
