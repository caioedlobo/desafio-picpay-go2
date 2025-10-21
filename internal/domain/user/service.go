package user

import (
	"context"
	"desafio-picpay-go2/internal/common/dto"
	"desafio-picpay-go2/internal/domain/user/value_object"
	"desafio-picpay-go2/pkg/fault"
	"desafio-picpay-go2/pkg/strutil"
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
	userExists, err := s.repo.FindByEmail(ctx, input.Email)
	if err != nil {
		s.log.Error("error finding user by email", "user", input.Email)
		return nil, err
	}
	if userExists != nil {
		s.log.Error("user already exists", "user", input.Email)
		return nil, fault.NewBadRequest("user already exists")
	}
	name, err := value_object.NewName(input.Name)
	if err != nil {
		s.log.Error("error creating new user", "name", input.Name, "err", err)
		return nil, err
	}
	docNumber, err := value_object.NewDocumentNumber(input.DocumentNumber)
	if err != nil {
		s.log.Error("error creating new user", "docNumber", input.DocumentNumber, "err", err)
		return nil, err
	}
	docType, err := value_object.NewDocumentType(input.DocumentType)
	if err != nil {
		s.log.Error("error creating new user", "docType", input.DocumentType, "err", err)
		return nil, err
	}
	email, err := value_object.NewEmail(input.Email)
	if err != nil {
		s.log.Error("error creating new user", "email", input.Email, "err", err)
		return nil, err
	}
	password, err := value_object.NewPassword(input.Password)
	if err != nil {
		s.log.Error("error creating new user password", "err", err)
		return nil, err
	}
	u, err := NewUser(name, docNumber, docType, email, *password)
	if err != nil {
		s.log.Error("error creating new user", "err", err)
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
	s.log.Debug("user created successfully", "details", strutil.JSONStringify(userCreated))
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
		s.log.Debug("error trying to match password", "err", err)
		return nil, err
	}
	if !match {
		s.log.Debug("password does not match", "user", input.Email)
		return nil, ErrUserNotFound
	}
	tkn, _, err := token.Gen(s.secretKey, foundUser.ID, s.accessTokenDuration)
	if err != nil {
		s.log.Debug("error generating token", "err", err)
		return nil, err
	}

	return &dto.LoginResponse{AccessToken: tkn}, nil
}
