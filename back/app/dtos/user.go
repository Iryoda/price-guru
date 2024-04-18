package dtos

import "github.com/iryoda/price-guru/app/entities"

type CreateUserResponse struct {
	User  entities.User `json:"user"`
	Token string        `json:"token"`
}
