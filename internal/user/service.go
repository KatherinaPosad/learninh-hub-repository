package user

import (
	"context"
	"learning-hub/auth"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *Repository
	jwtSecret  string
}

func NewService(repository *Repository, secret string) *Service {
	return &Service{
		repository: repository,
		jwtSecret:  secret,
	}
}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrInvalidEmail
	}
	return nil
}

func (s *Service) Register(ctx context.Context, request RegisterRequest) (UserResponse, error) {
	if request.Login == "" && request.Email == "" {
		return UserResponse{}, ErrLoginOrEmailRequired
	}

	if request.Email != "" {
		err := validateEmail(request.Email)
		if err != nil {
			return UserResponse{}, err
		}
	}

	if request.Password == "" {
		return UserResponse{}, ErrPasswordRequired
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

func (s *Service) Login(ctx context.Context, request LoginRequest) (LoginResponse, error) {

	if request.Login == "" && request.Email == "" { //нельзя указать ничего
		return LoginResponse{}, ErrLoginOrEmailRequired
	}

	if request.Login != "" && request.Email != "" { // нельзя указать оба
		return LoginResponse{}, ErrLoginOrEmailRequired
	}

	if request.Password == "" {
		return LoginResponse{}, ErrPasswordRequired
	}

	var user User
	var err error

	if request.Login != "" {
		user, err = s.repository.GetUserByLogin(ctx, request.Login)
	} else {
		user, err = s.repository.GetUserByEmail(ctx, request.Email)
	}

	if err != nil {
		return LoginResponse{}, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(request.Password),
	)

	if err != nil {
		return LoginResponse{}, ErrInvalidCredentials
	}

	roles := []string{user.Role}

	token, err := auth.GenerateJWT(user.ID, roles, s.jwtSecret)
	if err != nil {
		return LoginResponse{}, err
	}

	userResponse := UserResponse{
		ID:    user.ID,
		Login: user.Login,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	response := LoginResponse{
		User:  userResponse,
		Token: token,
	}

	return response, nil
}
