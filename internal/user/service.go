package user

import (
	"context"
	"errors"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email")
	}
	return nil
}

func (s *Service) Register(ctx context.Context, request RegisterRequest) (UserResponse, error) {
	if request.Login == "" && request.Email == "" {
		return UserResponse{}, errors.New("Login or must be provided")
	}

	if request.Email != "" {
		err := validateEmail(request.Email)
		if err != nil {
			return UserResponse{}, err
		}
	}

	if request.Password == "" {
		return UserResponse{}, errors.New("Password must be not empty")
	}

	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return UserResponse{}, err
	}

	user := User{
		Login:        request.Login,
		Name:         "",
		Email:        request.Email,
		PasswordHash: string(hashPassword),
		Role:         "user",
	}

	err = s.repository.CreateUser(ctx, &user)
	if err != nil {
		return UserResponse{}, err
	}

	response := UserResponse{
		ID:    user.ID,
		Login: user.Login,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
	return response, nil
}

func (s *Service) Login(ctx context.Context, request LoginRequest) (User, error) {
	var user User
	var err error

	if request.Login == "" && request.Email == "" { //нельзя указать ничего
		return User{}, errors.New("Login and Email must be provided")
	}

	if request.Login != "" && request.Email != "" { // нельзя указать оба
		return User{}, errors.New("provide login or email")
	}

	if request.Password == "" {
		return User{}, errors.New("Password must be not empty")
	}

	if request.Login != "" {
		user, err = s.repository.GetUserByLogin(ctx, request.Login)
	} else {
		user, err = s.repository.GetUserByEmail(ctx, request.Email)
	}

	if err != nil {
		if request.Login != "" {
			return User{}, errors.New("invalid login or password")
		}
		return User{}, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(request.Password),
	)

	if err != nil {
		if request.Login != "" {
			return User{}, errors.New("invalid login or password")
		}
		return User{}, errors.New("invalid email or password")
	}

	return user, nil
}
