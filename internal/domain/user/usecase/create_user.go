package usecase

import (
	"context"
	"desafio-picpay-go2/internal/domain/user"
	"desafio-picpay-go2/internal/domain/user/value_object"
	"errors"
)

type CreateUserRequest struct {
	Name           string `json:"name" validate:"required,lte=100"`
	DocumentNumber string `json:"document_number" validate:"required,gte=11,lte=14"`
	DocumentType   string `json:"document_type" validate:"required,gte=3,lte=4"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,lte=100"`
}

type CreateUserResponse struct {
	Id             string `json:"id" validate:"required"`
	Name           string `json:"name" validate:"required,lte=100"`
	DocumentNumber string `json:"document_number" validate:"required,gte=11,lte=14"`
	DocumentType   string `json:"document_type" validate:"required,gte=3,lte=4"`
	Email          string `json:"email" validate:"required,email"`
}

type UserRepository interface {
	Save(ctx context.Context, u *user.User) error
	FindByEmail(ctx context.Context, email string) (*user.User, error)
}

func (uc UserUseCase) Execute(ctx context.Context, input CreateUserRequest) (*CreateUserResponse, error) {
	if userExists, _ := uc.CreateUserRepo.FindByEmail(ctx, input.Email); userExists != nil {
		return nil, errors.New("email already registered")
	}

	name, err := value_object.NewName(input.Name)
	if err != nil {
		return nil, err
	}
	docNumber, err := value_object.NewDocumentNumber(input.DocumentNumber)
	if err != nil {
		return nil, err
	}
	docType, err := value_object.NewDocumentType(input.DocumentType)
	if err != nil {
		return nil, err
	}
	email, err := value_object.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}
	password, err := value_object.NewPassword(input.Password)
	if err != nil {
		return nil, err
	}
	u, err := user.NewUser(name, docNumber, docType, email, *password)
	if err != nil {
		return nil, err
	}

	if err = uc.CreateUserRepo.Save(ctx, u); err != nil {
		return nil, err
	}
	return &CreateUserResponse{
		Id:             u.ID.String(),
		Name:           u.Name.String(),
		DocumentNumber: u.DocumentNumber.String(),
		DocumentType:   u.DocumentType.String(),
		Email:          u.Email.String(),
	}, nil
}
