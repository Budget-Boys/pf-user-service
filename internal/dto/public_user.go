package dto

import (
	"user-service/internal/model"

	"github.com/google/uuid"
)

type PublicUser struct {
	ID      string `json:"id"`
	CpfCnpj string `json:"cpfcnpj"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

func ToPublicUser(user *model.User) PublicUser {
	return PublicUser{
		ID:      user.ID.String(),
		CpfCnpj: user.CPFCNPJ,
		Name:    user.Name,
		Email:   user.Email,
	}
}

func ToPublicUserFromUpdateInput(id uuid.UUID, input UserUpdateInput) PublicUser {
	return PublicUser{
		ID:      id.String(),
		CpfCnpj: input.CPFCNPJ,
		Name:    input.Name,
		Email:   input.Email,
	}
}
