package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/chacha20poly1305"
	"strings"
	"time"
)

func Gen(secretKey string, duration time.Duration) (string, *Claims, error) {
	if len(secretKey) != chacha20poly1305.KeySize {
		return "", nil, fmt.Errorf("invalid secret key")
	}

	claims, err := NewClaims(duration)
	if err != nil {
		return "", nil, fmt.Errorf("failed to create session claims: %w", err)
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", claims, fmt.Errorf("failed to sign token: %w", err)
	}

	return token, claims, nil
}

func Verify(secretKey, v string) (*Claims, error) {
	if strings.TrimSpace(v) == "" {
		return nil, fmt.Errorf("invalid token")
	}
	keyFunc := func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return []byte(secretKey), nil
	}
	token, err := jwt.ParseWithClaims(v, &Claims{}, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
