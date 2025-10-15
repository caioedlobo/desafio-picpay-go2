package model

import "time"

type User struct {
	ID             string    `db:"id"`
	Name           string    `db:"name"`
	DocumentNumber string    `db:"document_number"`
	DocumentType   string    `db:"document_type"`
	Email          string    `db:"email"`
	PasswordHash   []byte    `db:"password_hash"`
	CreatedAt      time.Time `db:"created_at"`
}
