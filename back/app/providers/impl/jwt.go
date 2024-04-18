package providers_impl

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iryoda/price-guru/app/providers"
)

type JWTProvider struct {
}

func (j JWTProvider) GenerateToken(claims providers.TokenClaims) (string, error) {
	c := providers.TokenClaims{
		UserId: claims.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 120)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	if token == nil {
		return "", errors.New("Error generating token")
	}

	signed, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return signed, nil
}

func (j JWTProvider) GetClaims(token string) (*providers.TokenClaims, error) {
	t, err := jwt.ParseWithClaims(token, &providers.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		secret := os.Getenv("JWT_SECRET")

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims := t.Claims.(*providers.TokenClaims)

	return claims, nil
}
