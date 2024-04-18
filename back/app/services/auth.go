package services

import (
	"errors"

	"github.com/iryoda/price-guru/app/dtos"
	"github.com/iryoda/price-guru/app/providers"
	"github.com/iryoda/price-guru/app/repositories"
)

type AuthService struct {
	UserRepository repositories.UserRepository
	HashProvider   providers.HashProvider
	TokenProvider  providers.TokenProvider
}

func (as AuthService) Login(params dtos.LoginParams) (*dtos.LoginResponse, error) {
	u, _ := as.UserRepository.FindByEmail(params.Email)

	if u == nil {
		return nil, errors.New("User not found")
	}

	err := as.HashProvider.Compare(params.Password, u.Password)

	if err != nil {
		return nil, errors.New("Invalid password")
	}

	claims := providers.TokenClaims{
		UserId: u.Id,
	}

	token, err := as.TokenProvider.GenerateToken(claims)

	if err != nil {
		return nil, err
	}

	return &dtos.LoginResponse{
		User:  *u,
		Token: token,
	}, nil
}
