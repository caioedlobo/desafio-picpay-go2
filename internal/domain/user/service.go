package user

import (
	"context"
	"desafio-picpay-go2/internal/common/dto"
	"desafio-picpay-go2/internal/domain/user/value_object"
	"desafio-picpay-go2/pkg/fault"
	"desafio-picpay-go2/pkg/token"
	"errors"
	"github.com/charmbracelet/log"
	"time"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrPasswordNotMatches = errors.New("password does not match")
)

type service struct {
	repo                UserRepository
	log                 *log.Logger
	secretKey           string
	accessTokenDuration time.Duration
}

func NewService(repo UserRepository, logger *log.Logger, secretKey string, accessTokenDuration time.Duration) *service {
	return &service{
		repo:                repo,
		log:                 logger,
		secretKey:           secretKey,
		accessTokenDuration: accessTokenDuration,
	}
}

func (s service) Register(ctx context.Context, input dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	s.log.Debug("trying to register a new user with",
		"name", input.Name,
		"email", input.Email)
	if userExists, _ := s.repo.FindByEmail(ctx, input.Email); userExists != nil {
		return nil, fault.NewBadRequest("user already exists3")
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
		s.log.Error("failed to insert user", "err", err)
		return nil, err
	}

	userCreated := &dto.CreateUserResponse{
		Id:             u.ID.String(),
		Name:           u.Name.String(),
		DocumentNumber: u.DocumentNumber.String(),
		DocumentType:   u.DocumentType.String(),
		Email:          u.Email.String(),
	}
	s.log.Debug("user created successfully: ", userCreated)
	return userCreated, nil
}

func (s service) Login(ctx context.Context, input dto.LoginRequest) (*dto.LoginResponse, error) {
	foundUser, err := s.repo.FindByEmail(ctx, input.Email)

	if err != nil {
		s.log.Debug("error finding user", "err", err, "email", input.Email)
		return nil, err
	}
	if foundUser == nil {
		s.log.Error(ErrUserNotFound, "email", input.Email)
		return nil, ErrUserNotFound
	}
	match, err := value_object.Matches(foundUser.PasswordHash, input.Password)
	if err != nil {
		s.log.Debug("error matching password", "err", err)
		return nil, err
	}
	if !match {
		s.log.Debug(ErrPasswordNotMatches)
		return nil, ErrPasswordNotMatches
	}
	tkn, _, err := token.Gen(s.secretKey, s.accessTokenDuration)
	if err != nil {
		s.log.Debug("error generating token", "err", err)
		return nil, err
	}

	return &dto.LoginResponse{AccessToken: tkn}, nil
}
