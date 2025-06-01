package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"time"
	"user-service/internal/logger"
	"user-service/internal/model"
	"user-service/internal/repository"
	"user-service/internal/utils"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const userCacheTTL = 10 * time.Minute

type UserService interface {
	Create(user *model.User) error
	GetByID(id string) (*model.PublicUser, error)
	GetAll() ([]model.User, error)
	Update(user *model.User) error
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

func (userService *userService) GetByID(id string) (*model.PublicUser, error) {
	cacheKey := fmt.Sprintf("user:%s", id)

	cached, err := userService.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var publicUser model.PublicUser
		if err := json.Unmarshal([]byte(cached), &publicUser); err == nil {
			logger.Log.Info("Cache hit for user", zap.String("userID", id))
			return &publicUser, nil
		}
	}

	user, err := userService.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	publicUser := model.ToPublicUser(user)

	data, _ := json.Marshal(publicUser)
	userService.redisClient.Set(ctx, cacheKey, data, userCacheTTL)
	logger.Log.Info("Cache miss for user", zap.String("userID", id))
	return &publicUser, nil
}

func (userService *userService) GetAll() ([]model.User, error) {
	return userService.userRepository.FindAll()
}

func (userService *userService) Update(user *model.User) error {
	if err := userService.userRepository.Update(user); err != nil {
		return err
	}
	userService.InvalidateCache(user.ID.String())
	return nil
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
