package user

import "errors"

var (
	ErrAccessTokenNotProvided = errors.New("access token not provided")
	ErrUserNotFound           = errors.New("user not found")
	ErrFailedInsertUser       = errors.New("failed to insert user")
	ErrUserAlreadyExists      = errors.New("user already exists")
)
