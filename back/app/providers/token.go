package providers

import "github.com/golang-jwt/jwt/v5"

type TokenProvider interface {
	GenerateToken(claims TokenClaims) (string, error)
	GetClaims(token string) (*TokenClaims, error)
}

type TokenClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}
