package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var userFieldNames = map[string]string{
	"Name":     "nome",
	"CPFCNPJ":  "CPF ou CNPJ",
	"Phone":    "telefone",
	"Email":    "e-mail",
	"Password": "senha",
}

func FormatValidationErrors(err error) string {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return "Erro de validação inesperado."
	}

	for _, fieldErr := range validationErrors {
		field := fieldErr.Field()
		fieldName := userFieldNames[field]
		if fieldName == "" {
			fieldName = field
		}

		switch fieldErr.Tag() {
		case "required":
			return fmt.Sprintf("O campo %s é obrigatório.", fieldName)
		case "email":
			return "O e-mail informado é inválido."
		case "min":
			return fmt.Sprintf("O campo %s deve ter no mínimo %s caracteres.", fieldName, fieldErr.Param())
		case "len":
			return fmt.Sprintf("O campo %s deve ter %s dígitos.", fieldName, fieldErr.Param())
		default:
			return fmt.Sprintf("O campo %s é inválido.", fieldName)
		}
	}

	return "Nenhum erro de validação específico encontrado."
}
