package dtos

import "github.com/iryoda/price-guru/app/entities"

type LoginParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	User  entities.User `json:"user"`
	Token string        `json:"token"`
}
