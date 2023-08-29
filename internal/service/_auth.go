package service

import (
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture/middleware"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	jwt2 "github.com/NicholasLiem/ModulAjar_Backend/utils/jwt"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"strings"
)

func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	var loginData dto.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Fail to login")
		return
	}

	/**
	Check Identifier By Email or Username
	*/

	var condition datastruct.UserModel
	if strings.Contains(loginData.Identifier, "@") {
		condition = datastruct.UserModel{Email: loginData.Identifier}
	} else {
		condition = datastruct.UserModel{Username: loginData.Identifier}
	}

	userData, err := datastruct.FindOneUser(condition)
	if err != nil {
		response.ErrorResponse(rw, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	if hasValidToken(r) {
		response.ErrorResponse(rw, http.StatusConflict, "Already logged in")
		return
	}

	/**
	Check Password
	*/
	err = userData.CheckPassword(loginData.Password)
	if err != nil {
		response.ErrorResponse(rw, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	/**
	Generate a new JWT Token
	*/
	token, err := jwt2.CreateJWT()
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Fail to generate JWT Token")
	}

	responseJson := map[string]string{
		"token": token.Token,
	}

	response.SuccessResponse(rw, http.StatusOK, "Successfully logged in", responseJson)
	return
}

func RegisterHandler(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	userID, err := VerifyUserId(id)
	if err != nil {
		response.ErrorResponse(rw, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var registerData dto.RegisterDTO
	err = json.NewDecoder(r.Body).Decode(&registerData)
	if err != nil {
		response.ErrorResponse(rw, http.StatusBadRequest, "Invalid request")
		return
	}

	if hasValidToken(r) {
		response.ErrorResponse(rw, http.StatusConflict, "Already logged in")
		return
	}

	if !isEmailValid(registerData.Email) {
		response.ErrorResponse(rw, http.StatusConflict, "Email is not valid")
		return
	}

	if strings.Contains(registerData.Username, "@") {
		response.ErrorResponse(rw, http.StatusConflict, "Username can't contain '@' symbol")
		return
	}

	userModel := datastruct.UserModel{
		UserID:   uint(userID),
		Username: registerData.Username,
		Email:    registerData.Email,
	}

	err = userModel.SetPassword(registerData.Password)
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Failed to set password")
		return
	}

	err = datastruct.CreateUser(&userModel)
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Failed to create user")
		return
	}

	/**
	Generate a new JWT Token
	*/
	token, err := jwt2.CreateJWT()
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Fail to generate JWT Token")
	}

	responseJson := map[string]interface{}{
		"user":  userModel,
		"token": token.Token,
	}

	response.SuccessResponse(rw, http.StatusOK, "Successfully registered in", responseJson)
	return
}

func hasValidToken(r *http.Request) bool {
	tokenStr := middleware.DecodeJWTToken(r)
	claims, err := jwt2.VerifyJWT(tokenStr)

	if err != nil || claims == nil {
		return false
	}

	return true
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
