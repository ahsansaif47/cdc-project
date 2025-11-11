package utils

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func PasswordValidator(fv validator.FieldLevel) bool {
	password := fv.Field().String()
	var (
		hasMinLen  = len(password) >= 8
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString(password)
		hasNumber  = regexp.MustCompile(`[0-9]`).MatchString(password)
		hasSpecial = regexp.MustCompile(`[!@#$%^&*()_\-=\[\]{}|;:'",.<>?/~\\]`).MatchString(password)
	)
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func ValidateRequest(val *validator.Validate, req any, customValErrors map[string]string) map[string]string {

	err := val.Struct(req)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	errors := make(map[string]string)

	for _, valErr := range validationErrors {
		fieldName := valErr.Field()

		// if you have a custom message defined
		if err_msg, exists := customValErrors[fieldName]; exists {
			errors[fieldName] = fmt.Sprintf("%s: %s", err_msg, err.Error())

		} else {
			// fallback to validator's default error message
			errors[fieldName] = valErr.Error()
		}
	}

	return errors
}
