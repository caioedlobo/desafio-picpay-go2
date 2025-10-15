package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
}

func NewClaims(duration time.Duration) (*Claims, error) {
	return &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}, nil
}

func (a *Claims) Valid() error {
	if time.Now().After(a.ExpiresAt.Time) {
		return errors.New("token has expired")
	}
	return nil
}
