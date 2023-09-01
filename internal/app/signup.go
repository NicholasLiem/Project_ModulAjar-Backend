package app

import (
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"net/http"
)

func (m *MicroserviceServer) Register(w http.ResponseWriter, r *http.Request) {
	var userModel datastruct.UserModel
	err := json.NewDecoder(r.Body).Decode(&userModel)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	token, err := m.authService.SignUp(userModel)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToRegister)
		return
	}

	responseMessage := map[string]string{
		"token": *token,
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulRegister, responseMessage)
	return
}
