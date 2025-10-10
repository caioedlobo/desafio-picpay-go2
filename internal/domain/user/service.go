package user

import (
	"context"
	"desafio-picpay-go2/internal/common/dto"
	"desafio-picpay-go2/internal/domain/user/value_object"
	"errors"
)

var ErrEmailAlreadyExists = errors.New("email already registered")

type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{repo: repo}
}

func (s Service) Register(ctx context.Context, input dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	if userExists, _ := s.repo.FindByEmail(ctx, input.Email); userExists != nil {
		return nil, ErrEmailAlreadyExists
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
	u, err := NewUser(name, docNumber, docType, email, *password)
	if err != nil {
		return nil, err
	}

	if err = s.repo.Save(ctx, u); err != nil {
		return nil, err
	}
	return &dto.CreateUserResponse{
		Id:             u.ID.String(),
		Name:           u.Name.String(),
		DocumentNumber: u.DocumentNumber.String(),
		DocumentType:   u.DocumentType.String(),
		Email:          u.Email.String(),
	}, nil
}
