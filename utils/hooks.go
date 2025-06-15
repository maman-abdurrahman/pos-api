package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidatorForm(error validator.FieldError) interface{} {
	errorValidateBody := make(map[string]string)
	fieldName := ToLower(error.Field())
	switch error.Tag() {
	case "required":
		errorValidateBody[fieldName] = fmt.Sprintf("Field %s is required", fieldName)
	case "min":
		errorValidateBody[fieldName] = fmt.Sprintf("Field %s is required, minimum %s characters", fieldName, error.Param())
	case "max":
		errorValidateBody[fieldName] = fmt.Sprintf("Field %s is required, maximum %s characters allowed", fieldName, error.Param())
	default:
		errorValidateBody[fieldName] = fmt.Sprintf("Failed on '%s' validation", error.Tag())
	}

	return errorValidateBody
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func Ucfirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func Ucwords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}
