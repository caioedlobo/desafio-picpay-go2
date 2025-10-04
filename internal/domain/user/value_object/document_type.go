package value_object

import (
	"errors"
	"strings"
)

const (
	CPF  DocumentType = "cpf"
	CNPJ DocumentType = "cnpj"
)

type DocumentType string

func NewDocumentType(value string) (DocumentType, error) {
	value = strings.ToLower(strings.TrimSpace(value))
	switch value {
	case string(CPF), string(CNPJ):
		return DocumentType(value), nil
	case "":
		return "", errors.New("document type should not be empty")
	default:
		return "", errors.New("document type is invalid")
	}
}

func (d DocumentType) IsCPF() bool {
	return d == CPF
}

func (d DocumentType) IsCNPJ() bool {
	return d == CNPJ
}

func (d DocumentType) String() string {
	return string(d)
}
