package error

import (
	"desafio-picpay-go2/pkg/fault"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ToFaultErrors(err error) []fault.FieldError {
	var fe []fault.FieldError
	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		for _, e := range validationErrs {
			fe = append(fe, fault.FieldError{
				Field:   e.Field(),
				Message: fmt.Sprintf("Field validation failed on the '%s' tag", e.Tag()),
			})
		}
	}
	return fe
}
