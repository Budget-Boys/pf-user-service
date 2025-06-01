package handler

import (
	"user-service/internal/model"
	"user-service/internal/service"
	validatorx "user-service/internal/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService service.AuthService
}

var validateAuth = validator.New()

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req model.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := validateAuth.Struct(req); err != nil {
		errors := validatorx.FormatValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}

	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
