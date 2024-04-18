package services

import (
	"errors"
	"fmt"

	"github.com/iryoda/price-guru/app/dtos"
	"github.com/iryoda/price-guru/app/entities"
	"github.com/iryoda/price-guru/app/providers"
	"github.com/iryoda/price-guru/app/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepository
	TokenProvider  providers.TokenProvider
	HashProvider   providers.HashProvider
}

func (us UserService) CreateUser(data entities.CreateUser) (*dtos.CreateUserResponse, error) {
	user, err := entities.NewUser(data)

	if err != nil {
		return nil, err
	}

	u, _ := us.UserRepository.FindByEmail(user.Email)

	if u != nil {
		return nil, errors.New("Email already taken")
	}

	hashedPassword, err := us.HashProvider.Hash(user.Password)

	if err != nil {
		return nil, errors.New("Error when parsing")
	}

	user.Password = hashedPassword

	user, err = us.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	token, err := us.TokenProvider.GenerateToken(providers.TokenClaims{
		UserId: user.Id,
	})

	fmt.Println("token", token)

	if err != nil {
		return nil, err

	}

	return &dtos.CreateUserResponse{User: *user, Token: token}, nil
}

func (us UserService) GetUserById(id string) (*entities.User, error) {
	user, err := us.UserRepository.FindById(id)

	if err != nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}
