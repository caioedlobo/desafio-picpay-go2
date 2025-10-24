package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type User struct {
	ID              string          `db:"id"`
	Name            string          `db:"name"`
	BalanceNumber   decimal.Decimal `db:"(balance).number"`
	BalanceCurrency string          `db:"(balance).currency_code"`
	DocumentNumber  string          `db:"document_number"`
	DocumentType    string          `db:"document_type"`
	Email           string          `db:"email"`
	PasswordHash    []byte          `db:"password_hash"`
	CreatedAt       time.Time       `db:"created_at"`
}
