package user

import (
	"context"
	"desafio-picpay-go2/internal/common/dto"
	"desafio-picpay-go2/internal/domain/user/value_object"
	"desafio-picpay-go2/internal/infra/http/middleware"
	"desafio-picpay-go2/pkg/fault"
	"desafio-picpay-go2/pkg/strutil"
	"desafio-picpay-go2/pkg/token"
	"errors"
	"github.com/charmbracelet/log"
	"time"
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

	name, err := value_object.NewName(input.Name)
	if err != nil {
		s.log.Error("error creating new user", "name", input.Name, "err", err)
		return nil, fault.NewBadRequest(err.Error())
	}
	docNumber, err := value_object.NewDocumentNumber(input.DocumentNumber)
	if err != nil {
		s.log.Error("error creating new user", "docNumber", input.DocumentNumber, "err", err)
		return nil, fault.NewBadRequest(err.Error())
	}
	docType, err := value_object.NewDocumentType(input.DocumentType)
	if err != nil {
		s.log.Error("error creating new user", "docType", input.DocumentType, "err", err)
		return nil, fault.NewBadRequest(err.Error())
	}
	email, err := value_object.NewEmail(input.Email)
	if err != nil {
		s.log.Error("error creating new user", "email", input.Email, "err", err)
		return nil, fault.NewBadRequest(err.Error())
	}
	password, err := value_object.NewPassword(input.Password)
	if err != nil {
		s.log.Error("error creating new user password", "err", err)
		return nil, fault.NewBadRequest(err.Error())
	}
	u, err := NewUser(name, docNumber, docType, email, *password)
	if err != nil {
		s.log.Error("error creating new user", "err", err)
		return nil, fault.NewInternalServerError("error creating new user")
	}

	if err = s.repo.Save(ctx, u); err != nil {
		s.log.Error(ErrFailedInsertUser, "err", err)
		switch {
		case errors.Is(err, ErrUserAlreadyExists):
			return nil, fault.NewConflict(ErrUserAlreadyExists.Error())
		default:
			return nil, fault.NewInternalServerError(ErrFailedInsertUser.Error())
		}
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
	s.log.Debug("trying to log user", "email", input.Email)

	foundUser, err := s.repo.FindByEmail(ctx, input.Email)
	if err != nil {
		s.log.Error("error logging user", "email", input.Email, "err", err)
		switch {
		case errors.Is(err, ErrUserNotFound):
			return nil, ErrUserNotFound
		default:
			return nil, errors.New("error finding user by email")
		}
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

func (s service) Get(ctx context.Context) (*dto.UserResponse, error) {
	s.log.Debug("trying to retrieve signed user")

	c, ok := ctx.Value(middleware.AuthKey{}).(*token.Claims)
	if !ok {
		s.log.Error("context does not contain auth key")
		return nil, ErrAccessTokenNotProvided
	}
	u, err := s.repo.FindByID(ctx, c.UserID)
	if err != nil {
		s.log.Error("error finding user", "id", c.UserID, "err", err)
		switch {
		case errors.Is(err, ErrUserNotFound):
			return nil, fault.NewBadRequest(ErrUserNotFound.Error())
		default:
			return nil, fault.NewInternalServerError("failed to find user by id")
		}
	}
	s.log.Debug("user retrieved successfully")

	return &dto.UserResponse{
		Email:          u.Email,
		Balance:        u.BalanceNumber,
		DocumentNumber: u.DocumentNumber,
		Name:           u.Name,
	}, nil
}
