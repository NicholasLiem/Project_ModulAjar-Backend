package auth

import (
	response "github.com/NicholasLiem/GoLang_Microservice/http"
	"net/http"
)

/**
TODO: Implement Login Handler
*/
func LoginHandler(rw http.ResponseWriter, r *http.Request) {

	response.SuccessResponse(rw, http.StatusOK, "Successfully logged in", "")
	return
}

/**
TODO: Implement Register Handler
*/
func RegisterHandler(rw http.ResponseWriter, r *http.Request) {

	response.SuccessResponse(rw, http.StatusOK, "Successfully registered in", "")
	return
}
