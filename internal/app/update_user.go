package app

import (
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	"github.com/NicholasLiem/ModulAjar_Backend/utils"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"github.com/gorilla/mux"
	"net/http"
)

func (m *MicroserviceServer) UpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	paramsUserID := params["user_id"]

	userID, err := utils.VerifyUserId(paramsUserID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseUserID)
		return
	}

	var updateUser dto.UpdateUserDTO
	err = json.NewDecoder(r.Body).Decode(&updateUser)
	updateUser.UserID = userID
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	updatedUser, err := m.userService.UpdateUser(updateUser)
	if err != nil || updatedUser == nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToUpdateUser)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulUserUpdate, updatedUser)
	return
}
