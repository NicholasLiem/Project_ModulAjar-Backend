package app

import (
	"github.com/NicholasLiem/ModulAjar_Backend/utils"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"github.com/gorilla/mux"
	"net/http"
)

func (m *MicroserviceServer) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramsUserID := params["user_id"]

	userID, err := utils.VerifyUserId(paramsUserID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseUserID)
		return
	}

	userData, err := m.userService.DeleteUser(userID)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToDeleteUser)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulUserDeletion, userData)
	return
}
