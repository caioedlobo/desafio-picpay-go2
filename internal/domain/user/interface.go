package user

import (
	"context"
	"desafio-picpay-go2/internal/common/dto"
	"desafio-picpay-go2/internal/infra/database/model"
)

type UserRepository interface {
	Save(context.Context, *User) error
	FindByEmail(context.Context, string) (*model.User, error)
}

type UserService interface {
	Register(context.Context, dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	Login(context.Context, dto.LoginRequest) (*dto.LoginResponse, error)
}
