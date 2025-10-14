package middleware

import (
	"context"
	"desafio-picpay-go2/pkg/token"
	"errors"
	"net/http"
	"strings"
)

type AuthKey struct{}

type middleware struct {
	secretKey string
}

func NewWithAuth(secretKey string) *middleware {
	return &middleware{secretKey: secretKey}
}

func (m *middleware) WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		tkn, err := validateAccessToken(accessToken)
		if err != nil {
			//TODO: Lançar erro
			return
		}

		claims, err := token.Verify(m.secretKey, tkn)
		if err != nil {
			if strings.Contains(err.Error(), "token has expired") {
				//TODO: Lançar erro
				return
			}
			//TODO: Lançar erro
			return
		}
		ctx := context.WithValue(r.Context(), AuthKey{}, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateAccessToken(header string) (string, error) {
	if strings.TrimSpace(header) == "" {
		return "", errors.New("missing authentication token")
	}

	parts := strings.Fields(header)

	if len(parts) != 2 {
		return "", errors.New("invalid auth token format")
	}

	if parts[0] != "Bearer" {
		return "", errors.New("invalid auth token format")
	}

	return parts[1], nil
}
