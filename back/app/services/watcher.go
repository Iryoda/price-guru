package services

import (
	"errors"

	"github.com/iryoda/price-guru/app/entities"
	"github.com/iryoda/price-guru/app/repositories"
)

type WatcherService struct {
	WatcherRepository repositories.WatcherRepository
}

func (ws WatcherService) CreateWatcher(data entities.CreateWatcher) (*entities.Watcher, error) {
	watcher, err := entities.NewWatcher(data)

	if err != nil {
		return nil, err
	}

	watcher, err = ws.WatcherRepository.Create(watcher)

	if err != nil {
		return nil, err
	}

	return watcher, nil
}

func (ws WatcherService) GetWatchersByUserId(id string) (*[]entities.Watcher, error) {
	watchers, err := ws.WatcherRepository.FindAllByUserId(id)

	if err != nil {
		return nil, errors.New("No watchers founded")
	}

	return watchers, nil
}

func (ws WatcherService) DeleteById(id string, userId string) error {
	watcher, err := ws.WatcherRepository.FindById(id)

	if err != nil {
		return errors.New("Watcher not found")
	}

	if watcher.UserId == userId {
		return errors.New("User does not have permission")
	}

	return nil
}
