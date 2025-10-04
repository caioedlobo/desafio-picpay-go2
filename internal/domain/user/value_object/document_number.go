package value_object

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrEmptyDocumentNumber   = errors.New("document number should not be empty")
	ErrInvalidDocumentNumber = errors.New("document number is invalid")
)

var DocumentNumberRX = regexp.MustCompile(`^\d{11}$|^\d{14}$`)

type DocumentNumber string

func NewDocumentNumber(value string) (DocumentNumber, error) {
	value = normalize(value)

	if value == "" {
		return "", ErrEmptyDocumentNumber
	}
	if matched := DocumentNumberRX.MatchString(value); !matched {
		return "", ErrInvalidDocumentNumber
	}
	return DocumentNumber(value), nil
}

func normalize(value string) string {
	replacer := strings.NewReplacer(".", "", "-", "", "/", "", " ", "")
	return replacer.Replace(strings.TrimSpace(value))
}

func (d DocumentNumber) String() string {
	return string(d)
}
