package repository

import (
	"user-service/internal/dto"
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

func (userRepository *userRepository) Update(id string, input dto.UserUpdateInput) error {
	var existing model.User
	if err := userRepository.db.First(&existing, "id = ?", id).Error; err != nil {
		return err
	}

	updateData := map[string]interface{}{}

	if input.Name != "" && input.Name != existing.Name {
		updateData["name"] = input.Name
	}

	if input.CPFCNPJ != "" && input.CPFCNPJ != existing.CPFCNPJ {
		updateData["cpfcnpj"] = input.CPFCNPJ
	}

	if input.Email != "" && input.Email != existing.Email {
		updateData["email"] = input.Email
	}

	if input.Phone != "" && input.Phone != existing.Phone {
		updateData["phone"] = input.Phone
	}

	if input.Password != "" {
		updateData["password"] = input.Password
	}

	if len(updateData) == 0 {
		return nil
	}

	return userRepository.db.Model(&model.User{}).Where("id = ?", id).Updates(updateData).Error
}

func (userRepository *userRepository) Delete(id string) error {
	return userRepository.db.Delete(&model.User{}, "id = ?", id).Error
}
