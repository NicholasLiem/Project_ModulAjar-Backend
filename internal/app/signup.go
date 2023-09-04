package app

import (
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/utils"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"github.com/gorilla/securecookie"
	"net/http"
	"os"
	"time"
)

func (m *MicroserviceServer) Register(w http.ResponseWriter, r *http.Request) {
	var userModel datastruct.UserModel
	err := json.NewDecoder(r.Body).Decode(&userModel)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	/**
	Parsing Session Data from Context
	*/
	userSession, err := utils.ParseCookie(r)
	if err == nil || userSession != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.AlreadyLoggedIn)
		return
	}

	if userModel.Email == "" || userModel.FirstName == "" || userModel.LastName == "" || userModel.Password == "" {
		response.ErrorResponse(w, http.StatusBadRequest, messages.AllFieldMustBeFilled)
		return
	}

	userData, err := m.authService.SignUp(userModel)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToRegister)
		return
	}

	sessionId, err := m.sessionService.CreateUserSession(*userData)
	if err != nil {
		response.ErrorResponse(w, http.StatusForbidden, messages.AlreadyLoggedIn)
		return
	}

	var (
		hashKey  = []byte(os.Getenv("HASH_KEY"))
		blockKey = []byte(os.Getenv("BLOCK_KEY"))
		s        = securecookie.New(hashKey, blockKey)
	)

	encoded, err := s.Encode("sessionId", *sessionId)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToEncodeCookie)
		return
	}

	responseCookie := http.Cookie{
		Name:     "sessionId",
		Value:    encoded,
		Expires:  time.Now().Add(1 * time.Hour),
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, &responseCookie)

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulRegister, nil)
	return
}
