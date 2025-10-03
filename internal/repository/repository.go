package repository

import (
	"desafio-picpay-go2/internal/domain/user"
	userrepo "desafio-picpay-go2/internal/repository/user"
)

type User interface {
	CreateUser(*user.User) error
}

type Repository struct {
	User
}

func NewRepository() *Repository {
	return &Repository{User: userrepo.New()}
}
