package jwt

import (
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type JWTToken struct {
	Token string `json:"token"`
}

func CreateJWT(payload datastruct.JWTPayload) (JWTToken, error) {

	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString(signingKey)

	return JWTToken{Token: tokenString}, err // Fixed token structure and variable name
}

func VerifyJWT(tokenStr string) (*datastruct.JWTPayload, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*datastruct.JWTPayload)
	if !ok {
		return nil, err // Handle the type assertion failure
	}

	return claims, nil // Return claims directly
}
