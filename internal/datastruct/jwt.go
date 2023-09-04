package datastruct

import "github.com/golang-jwt/jwt/v5"

type JWTPayload struct {
	UserId string `json:"user_id"`
	Role   Role   `json:"role"`
	jwt.RegisteredClaims
}
