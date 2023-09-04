package app

import (
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"net/http"
)

func (m *MicroserviceServer) Logout(w http.ResponseWriter, r *http.Request) {
	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulLogout, nil)
	return
}
