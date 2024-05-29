package handler

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

// initializes the validator with a custom validation rule for the cart state
func InitValidator() {
	validate = validator.New()
	err := validate.RegisterValidation("customState", validState)
	if err != nil {
		panic(err)
	}
}
