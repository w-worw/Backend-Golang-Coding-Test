package services

import (
	"7-solutions/dtos"
	"7-solutions/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(userDto *dtos.UserRegister) (*dtos.UserResponse, error)
	GetUserByID(id int) (*dtos.UserResponse, error)
	GetAllUsers() ([]dtos.UserResponse, error)
	UpdateUser(id int, userDto *dtos.UserUpdate) (*dtos.UserResponse, error)
	DeleteUser(id int) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(userDto *dtos.UserRegister) (*dtos.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userDto.Password = string(hashedPassword)

	user, err := s.userRepository.CreateUser(userDto)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByID(id int) (*dtos.UserResponse, error) {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetAllUsers() ([]dtos.UserResponse, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) UpdateUser(id int, userDto *dtos.UserUpdate) (*dtos.UserResponse, error) {
	user, err := s.userRepository.UpdateUser(id, userDto)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id int) error {
	err := s.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
