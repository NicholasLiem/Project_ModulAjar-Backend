package app

import (
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"net/http"
)

func (m *MicroserviceServer) Login(w http.ResponseWriter, r *http.Request) {
	var loginDTO dto.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&loginDTO)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	token, err := m.authService.SignIn(loginDTO)
	if err != nil {
		response.ErrorResponse(w, http.StatusUnauthorized, messages.UnsuccessfulLogin)
		return
	}

	responseMessage := map[string]string{
		"token": *token,
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulLogin, responseMessage)
	return
}
