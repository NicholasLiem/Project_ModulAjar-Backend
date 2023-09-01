package app

import (
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	"github.com/NicholasLiem/ModulAjar_Backend/utils"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	"github.com/gorilla/mux"
	"net/http"
)

func (m *MicroserviceServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	userID, err := utils.VerifyUserId(id)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var newUser dto.CreateUserDTO
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Fail to decode user")
		return
	}

	userModel := datastruct.UserModel{
		UserID:   userID,
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
	}

	err = m.userService.CreateUser(userModel)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Fail to create user"+err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, "User created", userModel)
	return
}
