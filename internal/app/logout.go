package app

import (
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"net/http"
)

func (m *MicroserviceServer) Logout(w http.ResponseWriter, r *http.Request) {

	sessionIdCookie, err := r.Cookie("sessionId")
	if err != nil {
		response.ErrorResponse(w, http.StatusForbidden, messages.SessionExpired)
		return
	}

	sessionId := sessionIdCookie.Value

	err = m.sessionService.InvalidateUserSession(sessionId)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.InvalidRequestData)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulLogout, nil)
	return
}
