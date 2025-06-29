package services

import (
	"7-solutions/dtos"
	"7-solutions/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(userDto *dtos.UserRegister) error
	AuthenticateUser(input *dtos.UserAuthenticate) (*string, error)
}

type authService struct {
	authRepository repositories.AuthRepository
}

func NewAuthService(authRepository repositories.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}

func (s *authService) RegisterUser(userDto *dtos.UserRegister) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userDto.Password = string(hashedPassword)

	err = s.authRepository.RegisterUser(userDto)
	if err != nil {
		return err
	}
	return nil
}

func (s *authService) AuthenticateUser(input *dtos.UserAuthenticate) (*string, error) {
	token, err := s.authRepository.AuthenticateUser(input)
	if err != nil {
		return nil, err
	}
	return token, nil
}

