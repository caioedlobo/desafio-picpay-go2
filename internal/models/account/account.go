package account

import (
	"desafio-picpay-go2/internal/models/account/value_object"
	"errors"
	"github.com/bojanz/currency"
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID             uuid.UUID
	FullName       string
	DocumentNumber string
	DocumentType   value_object.DocumentType
	Email          value_object.Email
	Password       value_object.Password
	Balance        currency.Amount
	CreatedAt      time.Time
}

func NewAccount(name, documentNumber string, documentType value_object.DocumentType, email value_object.Email, password value_object.Password) (*Account, error) {
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

	return &Account{
		ID:             uuid.New(),
		FullName:       name,
		DocumentNumber: documentNumber,
		DocumentType:   documentType,
		Email:          email,
		Password:       password,
		Balance:        balance,
	}, nil
}
