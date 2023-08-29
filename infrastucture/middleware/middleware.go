package middleware

import (
	"fmt"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	jwt2 "github.com/NicholasLiem/ModulAjar_Backend/utils/jwt"
	"net/http"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		tokenStr := DecodeJWTToken(r)
		if tokenStr == "" {
			response.ErrorResponse(rw, http.StatusUnauthorized, "Fail to verify JWT Token")
			return
		}

		claims, err := jwt2.VerifyJWT(tokenStr)

		if err != nil {
			response.ErrorResponse(rw, http.StatusUnauthorized, "Fail to verify JWT token: "+err.Error())
			return
		}

		fmt.Println(claims)
		next.ServeHTTP(rw, r)
	})
}

func DecodeJWTToken(r *http.Request) string {
	tokenStr := r.Header.Get("Authorization")

	if len(tokenStr) == 0 {
		return ""
	}

	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
	return tokenStr
}
