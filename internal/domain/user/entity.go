package user

import (
	"desafio-picpay-go2/internal/domain/user/value_object"
	"errors"
	"github.com/bojanz/currency"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID             uuid.UUID
	Name           value_object.Name
	DocumentNumber value_object.DocumentNumber
	DocumentType   value_object.DocumentType
	Email          value_object.Email
	Password       value_object.Password
	Balance        currency.Amount
	CreatedAt      time.Time
}

func NewUser(name value_object.Name, documentNumber value_object.DocumentNumber, documentType value_object.DocumentType, email value_object.Email, password value_object.Password) (*User, error) {
	if name == "" {
		return nil, errors.New("nome não pode ser vazio")
	}

	if documentNumber == "" {
		return nil, errors.New("número do documento não pode ser vazio")
	}

	if password.GetPlaintext() == nil {
		return nil, errors.New("senha não pode ser vazia")
	}
	balance, err := currency.NewAmount("0.0", "BRL")
	if err != nil {
		return nil, err
	}

	return &User{
		ID:             uuid.New(),
		Name:           name,
		DocumentNumber: documentNumber,
		DocumentType:   documentType,
		Email:          email,
		Password:       password,
		Balance:        balance,
	}, nil
}
