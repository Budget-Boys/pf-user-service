package dto

type UserUpdateInput struct {
	Name     string `json:"name" validate:"min=3"`
	CPFCNPJ  string `json:"cpfcnpj,omitempty" validate:"omitempty,len=11|len=14"`
	Phone    string `json:"phone"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"min=6"`
}
