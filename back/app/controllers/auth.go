package controllers

import (
	"net/http"

	"github.com/iryoda/price-guru/app/dtos"
	"github.com/iryoda/price-guru/app/entities"
	"github.com/iryoda/price-guru/app/services"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthService *services.AuthService
	UserService *services.UserService
}

func (ac AuthController) Login(c echo.Context) error {
	params := new(dtos.LoginParams)

	if err := c.Bind(params); err != nil {
		return err
	}

	result, err := ac.AuthService.Login(*params)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return c.JSON(http.StatusOK, result)
}

func (ac AuthController) CreateUser(c echo.Context) error {
	u := new(entities.CreateUser)

	if err := c.Bind(u); err != nil {
		return err
	}

	user, err := ac.UserService.CreateUser(*u)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return c.JSON(http.StatusCreated, user)
}
