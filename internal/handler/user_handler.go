package handler

import (
	"user-service/internal/model"
	"user-service/internal/service"

	validatorx "user-service/internal/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{
	userService service.UserService
}

var validate = validator.New()

func NewUserHandler(userService service.UserService) *UserHandler{
	return &UserHandler{userService}
}

func (userHandler *UserHandler) Create(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := validate.Struct(user); err != nil {
		errors := validatorx.FormatValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	if err := userHandler.userService.Create(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (userHandler *UserHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := userHandler.userService.GetByID(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func (userHandler *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := userHandler.userService.GetAll()

	if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(users)
}

func (userHandler *UserHandler) Delete(c *fiber.Ctx) error {
    id := c.Params("id")
    if err := userHandler.userService.Delete(id); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.SendStatus(fiber.StatusNoContent)
}