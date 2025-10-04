package value_object

import "errors"

type Name string

var ErrEmptyName = errors.New("name should not be empty")

func NewName(value string) (Name, error) {
	if value == "" {
		return "", ErrEmptyName
	}
	return Name(value), nil
}

func (n Name) String() string {
	return string(n)
}
