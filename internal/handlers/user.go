package handlers

import (
	"encoding/json"
	"errors"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/service"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		UserService: userService,
	}
}

func (uh *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	userID, err := VerifyUserId(id)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var newUser dto.CreateUserDTO
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Fail to decode user")
		return
	}

	userModel := datastruct.UserModel{
		UserID:   uint(userID),
		Username: newUser.Username,
		Email:    newUser.Email,
	}

	err = uh.UserService.CreateUser(userModel)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Fail to create user")
		return
	}

	response.SuccessResponse(w, http.StatusOK, "User created", userModel)
	return
}

func VerifyUserId(UserID string) (uint64, error) {
	userID, err := strconv.ParseUint(UserID, 10, 64)
	if err != nil {
		return 0, errors.New("cannot parse id")
	}
	return userID, nil
}
