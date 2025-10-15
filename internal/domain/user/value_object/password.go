package value_object

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	plaintext *string
	hash      []byte
}

func NewPassword(plaintext string) (*Password, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), 12)
	if err != nil {
		return nil, err
	}
	return &Password{hash: hash, plaintext: &plaintext}, nil
}

func (p *Password) GetHash() []byte {
	return p.hash
}

func (p *Password) GetPlaintext() *string {
	return p.plaintext
}

func Matches(hash []byte, plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
