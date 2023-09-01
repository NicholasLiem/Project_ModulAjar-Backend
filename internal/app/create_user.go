package app

import (
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	"github.com/NicholasLiem/ModulAjar_Backend/utils"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"github.com/gorilla/mux"
	"net/http"
)

func (m *MicroserviceServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	userID, err := utils.VerifyUserId(id)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseUserID)
		return
	}

	var newUser dto.CreateUserDTO
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
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
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToCreateUser+err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulUserCreation, userModel)
	return
}
