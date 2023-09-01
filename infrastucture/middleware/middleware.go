package middleware

import (
	"context"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	jwt2 "github.com/NicholasLiem/ModulAjar_Backend/utils/jwt"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		tokenStr := jwt2.DecodeJWTToken(r)
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
