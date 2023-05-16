package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("Key: '%s' Error: Validation for '%s' failed - '%s' tag id", v.Namespace(), v.Field(), v.Tag())
}
