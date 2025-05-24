package service

import (
	"user-service/internal/model"
	"user-service/internal/repository"
	"user-service/internal/utils"

	"github.com/google/uuid"
)

type UserService interface {
    Create(user *model.User) error
    GetByID(id string) (*model.User, error)
    GetAll() ([]model.User, error)
    Update(user *model.User) error
    Delete(id string) error
}

type userService struct {
    userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
    return &userService{userRepository}
}

func (userService *userService) Create(user *model.User) error {
    user.ID = uuid.New()
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return nil
	}

	user.Password = hashedPassword

    return userService.userRepository.Create(user)
}

func (userService *userService) GetByID(id string) (*model.User, error) {
    return userService.userRepository.FindByID(id)
}

func (userService *userService) GetAll() ([]model.User, error) {
    return userService.userRepository.FindAll()
}

func (userService *userService) Update(user *model.User) error {
    return userService.userRepository.Update(user)
}

func (userService *userService) Delete(id string) error {
    return userService.userRepository.Delete(id)
}
