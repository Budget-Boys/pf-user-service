package dto

type UserUpdateInput struct {
	Name     string `json:"name" validate:"required,min=3"`
	CPFCNPJ  string `json:"cpfcnpj" validate:"required,len=11|len=14"`
	Phone    string `json:"phone"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
