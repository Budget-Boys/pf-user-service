package repository

import (
	"user-service/internal/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByID(id string) (*model.User, error)
	FindAll() ([]model.User, error)
	Update(user *model.User) error
	Delete(id string) error
}
