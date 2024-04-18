package controllers

import (
	"net/http"

	"github.com/iryoda/price-guru/app/entities"
	"github.com/iryoda/price-guru/app/services"
	"github.com/labstack/echo/v4"
)

type WatcherController struct {
	WatcherService *services.WatcherService
}

func (wc WatcherController) CreateWatcher(c echo.Context) error {
	userId := c.Get("User").(string)

	w := new(entities.CreateWatcher)

	if err := c.Bind(w); err != nil {
		return err
	}

	w.UserId = userId

	watcher, err := wc.WatcherService.CreateWatcher(*w)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return c.JSON(http.StatusCreated, watcher)
}

func (wc WatcherController) GetWatchersByUserId(c echo.Context) error {
	userId := c.Get("User").(string)

	watcher, err := wc.WatcherService.GetWatchersByUserId(userId)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return c.JSON(http.StatusOK, watcher)
}

func (wc WatcherController) DeleteById(c echo.Context) error {
	id := c.Param("id")
	userId := c.Get("User").(string)

	err := wc.WatcherService.DeleteById(id, userId)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return c.NoContent(http.StatusNoContent)
}
