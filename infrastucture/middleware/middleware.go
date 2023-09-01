package middleware

import (
	"context"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	jwt2 "github.com/NicholasLiem/ModulAjar_Backend/utils/jwt"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"net/http"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		tokenStr := DecodeJWTToken(r)
		if tokenStr == "" {
			response.ErrorResponse(rw, http.StatusUnauthorized, messages.JWTClaimError)
			return
		}

		claims, err := jwt2.VerifyJWT(tokenStr)

		if err != nil {
			response.ErrorResponse(rw, http.StatusUnauthorized, messages.JWTClaimError)
			return
		}

		ctx := context.WithValue(r.Context(), "jwtClaims", claims)
		r = r.WithContext(ctx)

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
