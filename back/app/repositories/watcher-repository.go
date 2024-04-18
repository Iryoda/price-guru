package repositories

import "github.com/iryoda/price-guru/app/entities"

type WatcherRepository interface {
	Create(watcher *entities.Watcher) (*entities.Watcher, error)
	FindAllByUserId(id string) (*[]entities.Watcher, error)
	FindById(id string) (entities.Watcher, error)
	DeleteById(id string) error
}
