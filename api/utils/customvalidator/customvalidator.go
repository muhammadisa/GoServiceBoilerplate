package customvalidator

import (
	"gopkg.in/go-playground/validator.v9"
)

// CustomValidator struct
type CustomValidator struct {
	Validator *validator.Validate
}

// ICustomValidator interface
type ICustomValidator interface {
	Validate(i interface{}) error
}

// Validate validating an interface
func (cV CustomValidator) Validate(i interface{}) error {
	return cV.Validator.Struct(i)
}
