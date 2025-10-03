package userrepo

import "desafio-picpay-go2/internal/domain/user"

type UserRepository struct {
	userRepo *user.User
}

func New() *UserRepository {
	return &UserRepository{userRepo: nil}
}

func (u UserRepository) CreateUser(user *user.User) error {
	return nil
}
