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

func FormatValidationErrors(err error) map[string]string {
    errors := make(map[string]string)

    validationErrors, ok := err.(validator.ValidationErrors)
    if !ok {
        errors["erro"] = "Erro de validação inesperado"
        return errors
    }

    for _, fieldErr := range validationErrors {
        field := fieldErr.Field()
        fieldName := userFieldNames[field]
        if fieldName == "" {
            fieldName = field // fallback para o nome original
        }

        switch fieldErr.Tag() {
        case "required":
            errors[fieldName] = fmt.Sprintf("O campo %s é obrigatório", fieldName)
        case "email":
            errors[fieldName] = "O e-mail informado é inválido"
        case "min":
            errors[fieldName] = fmt.Sprintf("O campo %s deve ter no mínimo %s caracteres", fieldName, fieldErr.Param())
        case "len":
            errors[fieldName] = fmt.Sprintf("O campo %s deve ter %s dígitos", fieldName, fieldErr.Param())
        default:
            errors[fieldName] = fmt.Sprintf("O campo %s é inválido", fieldName)
        }
    }

    return errors
}
