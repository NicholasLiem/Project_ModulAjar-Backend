package app

import (
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	jwt2 "github.com/NicholasLiem/ModulAjar_Backend/utils/jwt"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"net/http"
	"strconv"
)

func (m *MicroserviceServer) Logout(w http.ResponseWriter, r *http.Request) {
	_, issuerUserID, err := jwt2.ParseUserIDClaim(r.Context())
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.JWTClaimError)
		return
	}

	err = m.sessionService.InvalidateUserSession(strconv.FormatUint(issuerUserID, 10))
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.InvalidRequestData)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulLogout, nil)
	return
}
