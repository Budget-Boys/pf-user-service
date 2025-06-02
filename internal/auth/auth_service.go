package auth

import (
	"errors"
	"time"
	"user-service/internal/repository"
	"user-service/internal/utils"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	Login(email, password string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type authService struct {
	userRepository repository.UserRepository
	jwtSecret      []byte
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string) AuthService {
	return &authService{
		userRepository: userRepo,
		jwtSecret:      []byte(jwtSecret),
	}
}

func (a *authService) Login(email, password string) (string, error) {
	user, err := a.userRepository.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	claims := jwt.MapClaims{
		"user_id":       user.ID.String(),
		"user_email":    user.Email,
		"user_name":     user.Name,
		"user_document": user.CPFCNPJ,
		"exp":           time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *authService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return a.jwtSecret, nil
	})
}
