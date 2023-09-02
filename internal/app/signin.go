package app

import (
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"net/http"
	"strconv"
)

func (m *MicroserviceServer) Login(w http.ResponseWriter, r *http.Request) {
	var loginDTO dto.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&loginDTO)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	_, err = m.authService.SignIn(loginDTO)
	if err != nil {
		response.ErrorResponse(w, http.StatusUnauthorized, messages.UnsuccessfulLogin)
		return
	}

	// implement session service
	// sementara dulu
	userModel := datastruct.UserModel{
		UserID:    10,
		FirstName: "Lol",
		LastName:  "Hai",
		Email:     "lol",
		Password:  "hehe",
		Role:      datastruct.ADMIN,
	}
	sessionId, err := m.sessionService.CreateUserSession(userModel)
	receivedData, err := m.sessionService.GetUserSession(*sessionId)
	responseMessage := map[string]string{
		"token":         *sessionId,
		"received_data": strconv.FormatBool(receivedData),
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulLogin, responseMessage)
	return
}
