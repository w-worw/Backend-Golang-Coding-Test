package services_test

import (
	"7-solutions/dtos"
	"7-solutions/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(userDto *dtos.UserRegister) (*dtos.UserResponse, error) {
	args := m.Called(userDto)
	return args.Get(0).(*dtos.UserResponse), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(id int) (*dtos.UserResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*dtos.UserResponse), args.Error(1)
}

func (m *MockUserRepository) GetAllUsers() ([]dtos.UserResponse, error) {
	args := m.Called()
	return args.Get(0).([]dtos.UserResponse), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(id int, userDto *dtos.UserUpdate) (*dtos.UserResponse, error) {
	args := m.Called(id, userDto)
	return args.Get(0).(*dtos.UserResponse), args.Error(1)
}

func (m *MockUserRepository) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) CountUsers() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func TestCreateUser_Success(t *testing.T) {
	repo := new(MockUserRepository)
	svc := services.NewUserService(repo)

	userInput := &dtos.UserRegister{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}
	userResponse := &dtos.UserResponse{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}

	repo.On("CreateUser", mock.AnythingOfType("*dtos.UserRegister")).Return(userResponse, nil)

	user, err := svc.CreateUser(userInput)
	assert.NoError(t, err)
	assert.Equal(t, userResponse.Name, user.Name)
	assert.Equal(t, userResponse.Email, user.Email)
	repo.AssertExpectations(t)
}

func TestCreateUser_HashPasswordError(t *testing.T) {
}

func TestGetUserByID_Success(t *testing.T) {
	repo := new(MockUserRepository)
	svc := services.NewUserService(repo)

	userID := 1
	userResponse := &dtos.UserResponse{
		ID:    userID,
		Name:  "Test User",
		Email: "test@example.com",
	}

	repo.On("GetUserByID", userID).Return(userResponse, nil)

	user, err := svc.GetUserByID(userID)
	assert.NoError(t, err)
	assert.Equal(t, userID, user.ID)
	repo.AssertExpectations(t)
}

func TestGetUserByID_NotFound(t *testing.T) {
	repo := new(MockUserRepository)
	svc := services.NewUserService(repo)

	userID := 999
	repo.On("GetUserByID", userID).Return((*dtos.UserResponse)(nil), errors.New("user not found"))

	user, err := svc.GetUserByID(userID)
	assert.Error(t, err)
	assert.Nil(t, user)
	repo.AssertExpectations(t)
}

func TestGetAllUsers_Success(t *testing.T) {
	repo := new(MockUserRepository)
	svc := services.NewUserService(repo)

	users := []dtos.UserResponse{
		{ID: 1, Name: "User1", Email: "user1@example.com"},
		{ID: 2, Name: "User2", Email: "user2@example.com"},
	}

	repo.On("GetAllUsers").Return(users, nil)

	result, err := svc.GetAllUsers()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	repo.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T) {
	repo := new(MockUserRepository)
	svc := services.NewUserService(repo)

	userID := 1
	updateData := &dtos.UserUpdate{
		Name:  "Updated Name",
		Email: "updated@example.com",
	}
	updatedUser := &dtos.UserResponse{
		ID:    userID,
		Name:  updateData.Name,
		Email: updateData.Email,
	}

	repo.On("UpdateUser", userID, updateData).Return(updatedUser, nil)

	result, err := svc.UpdateUser(userID, updateData)
	assert.NoError(t, err)
	assert.Equal(t, updateData.Name, result.Name)
	repo.AssertExpectations(t)
}

func TestDeleteUser_Success(t *testing.T) {
	repo := new(MockUserRepository)
	svc := services.NewUserService(repo)

	userID := 1
	repo.On("DeleteUser", userID).Return(nil)

	err := svc.DeleteUser(userID)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteUser_Failure(t *testing.T) {
	repo := new(MockUserRepository)
	svc := services.NewUserService(repo)

	userID := 2
	repo.On("DeleteUser", userID).Return(errors.New("delete failed"))

	err := svc.DeleteUser(userID)
	assert.Error(t, err)
	repo.AssertExpectations(t)
}
