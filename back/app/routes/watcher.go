package routes

import (
	"github.com/iryoda/price-guru/app/controllers"
	"github.com/labstack/echo/v4"
)

type WatcherRouter struct {
	WatcherController *controllers.WatcherController
}

func (w WatcherRouter) NewWatcherRouter(e *echo.Group) {
	g := e.Group("/watchers")

	g.POST("/create", w.WatcherController.CreateWatcher)
	g.GET("/all", w.WatcherController.GetWatchersByUserId)
}
