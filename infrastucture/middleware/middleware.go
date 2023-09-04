package middleware

import (
	"context"
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/utils"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func Middleware(next http.Handler, redisClient *redis.Client) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		sessionId, err := utils.ParseCookie(r)
		if err != nil {
			response.ErrorResponse(rw, http.StatusUnauthorized, messages.InvalidRequestData)
			return
		}

		sessionData, err := redisClient.Get(context.Background(), *sessionId).Bytes()
		if err != nil {
			response.ErrorResponse(rw, http.StatusInternalServerError, messages.SessionExpired)
			return
		}

		var sessionUserData datastruct.SessionUserClient
		if err = json.Unmarshal(sessionData, &sessionUserData); err != nil {
			response.ErrorResponse(rw, http.StatusInternalServerError, messages.FailToUnMarshalData)
			return
		}

		ctx := context.WithValue(r.Context(), "sessionData", sessionData)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
