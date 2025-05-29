package handler

import (
	"user-service/internal/dto"
	"user-service/internal/model"
	"user-service/internal/service"

	validatorx "user-service/internal/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService service.UserService
}

var validate = validator.New()

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (userHandler *UserHandler) Create(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return dto.RespondError(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := validate.Struct(user); err != nil {
		error := validatorx.FormatValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": error,
		})
	}

	if err := userHandler.userService.Create(&user); err != nil {
		return dto.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return dto.RespondSuccess(c, fiber.StatusCreated, "User created!", user)
}

func (userHandler *UserHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := userHandler.userService.GetByID(id)

	if err != nil {
		return dto.RespondError(c, fiber.StatusNotFound, "User not found")
	}

	return dto.RespondSuccess(c, fiber.StatusOK, "User retrieved!", user)
}

func (userHandler *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := userHandler.userService.GetAll()

	if err != nil {
		return dto.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return dto.RespondSuccess(c, fiber.StatusOK, "Users retrieved!", users)
}

func (userHandler *UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := userHandler.userService.Delete(id); err != nil {
		return dto.RespondError(c, fiber.StatusInternalServerError, err.Error())

	}
	return dto.RespondSuccess(c, fiber.StatusOK, "User deleted!", nil)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return dto.RespondError(c, fiber.StatusBadRequest, "Invalid Id")
	}

	var input dto.UserUpdateInput
	if err := c.BodyParser(&input); err != nil {
		return dto.RespondError(c, fiber.StatusBadRequest, "Invalid Input")
	}

	if err := validate.Struct(input); err != nil {
		error := validatorx.FormatValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": error,
		})
	}

	if err := h.userService.Update(id.String(), input); err != nil {
		return dto.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return dto.RespondSuccess(c, fiber.StatusOK, "User Updated", input)
}
