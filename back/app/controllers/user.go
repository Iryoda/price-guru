package controllers

import (
	"fmt"
	"net/http"

	"github.com/iryoda/price-guru/app/services"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService *services.UserService
}

func (uc UserController) Get(c echo.Context) error {
	id := c.Get("User").(string)

	fmt.Println("id", id)

	user, err := uc.UserService.GetUserById(id)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return c.JSON(http.StatusOK, user)
}
