package value_object

import (
	"errors"
	"regexp"
)

var (
	ErrEmptyEmail   = errors.New("email should not be empty")
	ErrInvalidEmail = errors.New("email is invalid")
)
var EmailRX = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

type Email string

func NewEmail(value string) (Email, error) {
	if value == "" {
		return "", ErrEmptyEmail
	}
	if ok := EmailRX.MatchString(value); !ok {
		return "", ErrInvalidEmail
	}
	return Email(value), nil
}

func (e Email) String() string {
	return string(e)
}
