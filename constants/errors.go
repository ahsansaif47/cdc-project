package constants

import "errors"

var CustomValidationErrors = map[string]string{
	"email":    "Email is incorrect",
	"password": "Passowrd is incorrect",
	"phone_no": "Phone No. is incorrect",
	"dob":      "Date of birth is incorrect",
}

var ErrUserAlreadyExists = errors.New("user already exists")
var ErrAccountDoesNotExist = errors.New("account does not exist")
var ErrFailedToHashPassword = errors.New("failed to hash password")
