package app

import (
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	jwt2 "github.com/NicholasLiem/ModulAjar_Backend/utils/jwt"
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

	_, isLoggedIn, err := jwt2.HasLoggedIn(r.Context())
	if !isLoggedIn {
		response.ErrorResponse(w, http.StatusBadRequest, messages.AlreadyLoggedIn)
		return
	}

	token, err := m.authService.SignIn(loginDTO)
	if err != nil {
		response.ErrorResponse(w, http.StatusUnauthorized, messages.UnsuccessfulLogin+err.Error())
		return
	}

	responseMessage := map[string]string{
		"token": *token,
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulLogin, responseMessage)
	return
}
