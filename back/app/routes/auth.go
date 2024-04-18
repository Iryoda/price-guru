package routes

import (
	"github.com/iryoda/price-guru/app/controllers"
	"github.com/labstack/echo/v4"
)

type AuthRouter struct {
	AuthController *controllers.AuthController
}

func (a AuthRouter) NewAuthRouter(e *echo.Echo) {
	e.POST("/auth/login", a.AuthController.Login)
	e.POST("/auth/signin", a.AuthController.CreateUser)
}
