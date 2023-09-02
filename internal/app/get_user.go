package app

import (
	"github.com/NicholasLiem/ModulAjar_Backend/utils"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	jwt2 "github.com/NicholasLiem/ModulAjar_Backend/utils/jwt"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"github.com/gorilla/mux"
	"net/http"
)

func (m *MicroserviceServer) GetUserData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramsUserID := params["user_id"]
	requestedUserID, err := utils.VerifyUserId(paramsUserID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseUserID)
		return
	}

	_, issuerUserID, err := jwt2.ParseUserIDClaim(r.Context())
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.JWTClaimError)
		return
	}

	userData, err := m.userService.GetUser(requestedUserID, uint(issuerUserID))
	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulUserObtain, userData)
	return
}
