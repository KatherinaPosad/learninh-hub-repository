package user

import "errors"

var (
	ErrLoginOrEmailRequired = errors.New("login or email must be provided")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrPasswordRequired     = errors.New("password must not be empty")
	ErrInvalidCredentials   = errors.New("invalid login or password")
	ErrUserAlreadyExists    = errors.New("user already exists")
)
