package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"user-service/internal/dto"
	"user-service/internal/logger"
	"user-service/internal/model"
	"user-service/internal/repository"
	"user-service/internal/utils"

	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const userCacheTTL = 10 * time.Minute

type UserService interface {
	Create(user *model.User) error
	GetByID(id string) (*dto.PublicUser, error)
	GetAll() ([]dto.PublicUser, error)
	Update(id string, input dto.UserUpdateInput) (*dto.PublicUser, error)
	Delete(id string) error
}

type userService struct {
	userRepository repository.UserRepository
	redisClient    *redis.Client
}

func NewUserService(userRepository repository.UserRepository, redisClient *redis.Client) UserService {
	return &userService{userRepository, redisClient}
}

var ctx = context.Background()

func (userService *userService) Create(user *model.User) error {
	user.ID = uuid.New()
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return userService.userRepository.Create(user)
}

func (userService *userService) GetByID(id string) (*dto.PublicUser, error) {
	cacheKey := fmt.Sprintf("user:%s", id)

	cached, err := userService.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var publicUser dto.PublicUser
		if err := json.Unmarshal([]byte(cached), &publicUser); err == nil {
			logger.Log.Info("Cache hit for user", zap.String("userID", id))
			return &publicUser, nil
		}
	}

	user, err := userService.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	publicUser := dto.ToPublicUser(user)

	data, _ := json.Marshal(publicUser)
	userService.redisClient.Set(ctx, cacheKey, data, userCacheTTL)
	logger.Log.Info("Cache miss for user", zap.String("userID", id))
	return &publicUser, nil
}

func (userService *userService) GetAll() ([]dto.PublicUser, error) {
	users, err := userService.userRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var publicUsers []dto.PublicUser

	for _, user := range users {
		publicUsers = append(publicUsers, dto.ToPublicUser(&user))
	}

	return publicUsers, nil
}

func (userService *userService) Update(id string, input dto.UserUpdateInput) (*dto.PublicUser, error) {
	existingUser, err := userService.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	updateData := make(map[string]interface{})

	if input.Name != "" && input.Name != existingUser.Name {
		updateData["name"] = input.Name
	}

	if input.CPFCNPJ != "" && input.CPFCNPJ != existingUser.CPFCNPJ {
		updateData["cpfcnpj"] = input.CPFCNPJ
	}

	if input.Email != "" && input.Email != existingUser.Email {
		updateData["email"] = input.Email
	}

	if input.Phone != "" && input.Phone != existingUser.Phone {
		updateData["phone"] = input.Phone
	}

	if input.Password != "" {
		updateData["password"], err = utils.HashPassword(input.Password)

		if err != nil {
			return nil, err
		}
	}

	if len(updateData) == 0 {
		return nil, errors.New("Nothing to update")
	}

	if err := userService.userRepository.Update(id, updateData); err != nil {
		return nil, err
	}

	user, err := userService.GetByID(id)

	if err != nil {
		return nil, err
	}

	userService.InvalidateCache(id)

	return user, nil
}

func (userService *userService) Delete(id string) error {
	if err := userService.userRepository.Delete(id); err != nil {
		return err
	}
	userService.InvalidateCache(id)
	return nil
}

func (userService *userService) InvalidateCache(id string) {
	cacheKey := fmt.Sprintf("user:%s", id)
	userService.redisClient.Del(ctx, cacheKey)
}
