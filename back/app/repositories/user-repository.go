package repositories

import "github.com/iryoda/price-guru/app/entities"

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	FindById(id string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
}
