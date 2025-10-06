package repository

import (
	"database/sql"
	"desafio-picpay-go2/internal/domain/user"
	"desafio-picpay-go2/internal/domain/user/usecase"
)

type Repository struct {
	userRepo usecase.UserRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{userRepo: user.NewRepository(db)}
}
