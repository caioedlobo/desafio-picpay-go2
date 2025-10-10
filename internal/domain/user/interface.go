package user

import (
	"context"
	"desafio-picpay-go2/internal/common/dto"
)

type UserRepository interface {
	Save(context.Context, *User) error
	FindByEmail(context.Context, string) (*User, error)
}

type UserService interface {
	Register(context.Context, dto.CreateUserRequest) (*dto.CreateUserResponse, error)
}
