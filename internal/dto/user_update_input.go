package dto

type UserUpdateInput struct {
	Name     string `json:"name" validate:"min=3"`
	CPFCNPJ  string `json:"cpfcnpj" validate:"len=11|len=14"`
	Phone    string `json:"phone"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=6"`
}
