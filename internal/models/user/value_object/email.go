package value_object

import (
	"errors"
	"regexp"
)

type Email string

var EmailRX = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func NewEmail(value string) (Email, error) {
	if ok, err := ValidEmail(value); !ok {
		return "", err
	}
	return Email(value), nil
}

func ValidEmail(value string) (bool, error) {
	if value == "" {
		return false, errors.New("email não pode ser vazio")
	}
	if ok := EmailRX.MatchString(value); !ok {
		return false, errors.New("email deve ser válido")
	}
	return true, nil
}
