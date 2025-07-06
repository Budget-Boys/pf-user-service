package repository

import (
	"user-service/internal/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) Create(user *model.User) error {
	return userRepository.db.Create(user).Error
}

func (userRepository *userRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	if err := userRepository.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := userRepository.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := userRepository.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Update(id string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	return r.db.Model(&model.User{}).Where("id = ?", id).Updates(updates).Error
}

func (userRepository *userRepository) Delete(id string) error {
	return userRepository.db.Delete(&model.User{}, "id = ?", id).Error
}
