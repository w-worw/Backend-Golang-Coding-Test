package services_test

import (
	"7-solutions/dtos"
	"7-solutions/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockAuthRepository struct {
	RegisterUserFunc     func(userDto *dtos.UserRegister) error
	AuthenticateUserFunc func(input *dtos.UserAuthenticate) (*string, error)
}

func (m *mockAuthRepository) RegisterUser(userDto *dtos.UserRegister) error {
	return m.RegisterUserFunc(userDto)
}

func (m *mockAuthRepository) AuthenticateUser(input *dtos.UserAuthenticate) (*string, error) {
	return m.AuthenticateUserFunc(input)
}

func TestRegisterUser_Success(t *testing.T) {
	mockRepo := &mockAuthRepository{
		RegisterUserFunc: func(userDto *dtos.UserRegister) error {
			return nil
		},
	}

	service := services.NewAuthService(mockRepo)

	err := service.RegisterUser(&dtos.UserRegister{
		Name:     "Test User",
		Email:    "test@user.com",
		Password: "password123",
	})

	assert.NoError(t, err)
}

func TestAuthenticateUser_Success(t *testing.T) {
	mockRepo := &mockAuthRepository{
		AuthenticateUserFunc: func(input *dtos.UserAuthenticate) (*string, error) {
			token := "mock_token"
			return &token, nil
		},
	}

	service := services.NewAuthService(mockRepo)

	token, err := service.AuthenticateUser(&dtos.UserAuthenticate{
		Email:    "test@user.com",
		Password: "password123",
	})

	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.Equal(t, "mock_token", *token)
}
