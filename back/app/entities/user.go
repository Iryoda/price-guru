package entities

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id       string `json:"id" bson:"_id,omitempty"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"-" bson:"password"`
}

type CreateUser struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

func NewUser(data CreateUser) (*User, error) {
	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		return nil, err
	}

	return &User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}, nil
}
