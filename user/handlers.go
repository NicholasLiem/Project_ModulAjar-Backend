package user

import (
	"encoding/json"
	"errors"
	response "github.com/NicholasLiem/ModulAjar_Backend/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateUserHandler(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	userID, err := VerifyUserId(id)
	if err != nil {
		response.ErrorResponse(rw, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var newUser CreateUserDTO
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Fail to decode user")
		return
	}

	userModel := UserModel{
		UserID:   uint(userID),
		Username: newUser.Username,
		Email:    newUser.Email,
	}

	err = userModel.SetPassword(newUser.Password)
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Failed to set password")
		return
	}
	err = CreateUser(&userModel)
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Failed to create user")
		return
	}

	response.SuccessResponse(rw, http.StatusCreated, "User created successfully", userModel)
	return
}

func FindUserByIdHandler(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	userID, err := VerifyUserId(id)
	if err != nil {
		response.ErrorResponse(rw, http.StatusBadRequest, "Invalid user ID")
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

	userID, err := VerifyUserId(id)
	if err != nil {
		response.ErrorResponse(rw, http.StatusBadRequest, "Invalid user ID")
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

	userID, err := VerifyUserId(id)
	if err != nil {
		response.ErrorResponse(rw, http.StatusBadRequest, "Invalid user ID")
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
		response.ErrorResponse(rw, http.StatusInternalServerError, "Fail to decode update data")
		return
	}

	err = foundUser.Update(updateData)
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Fail to update update data")
		return
	}

	response.SuccessResponse(rw, http.StatusOK, "Updated user data", updateData)
}

func VerifyUserId(UserID string) (uint64, error) {
	userID, err := strconv.ParseUint(UserID, 10, 64)
	if err != nil {
		return 0, errors.New("cannot parse id")
	}
	return userID, nil
}
