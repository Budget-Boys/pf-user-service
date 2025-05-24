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

func (userRepository *userRepository) FindAll() ([]model.User, error) {
    var users []model.User
    if err := userRepository.db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

func (userRepository *userRepository) Update(user *model.User) error {
    return userRepository.db.Save(user).Error
}

func (userRepository *userRepository) Delete(id string) error {
    return userRepository.db.Delete(&model.User{}, "id = ?", id).Error
}
