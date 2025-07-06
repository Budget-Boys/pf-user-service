package repository

import (
	"user-service/internal/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByID(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindAll() ([]model.User, error)
	Update(id string, input map[string]interface{}) error
	Delete(id string) error
}
