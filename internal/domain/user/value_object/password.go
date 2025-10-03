package value_object

import "golang.org/x/crypto/bcrypt"

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
