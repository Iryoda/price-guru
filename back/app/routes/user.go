package routes

import (
	"github.com/iryoda/price-guru/app/controllers"
	"github.com/labstack/echo/v4"
)

type UserRouter struct {
	UserController *controllers.UserController
}

func (r UserRouter) NewUserRouter(e *echo.Group) {
	g := e.Group("/users")

	g.GET("/info", r.UserController.Get)
}
