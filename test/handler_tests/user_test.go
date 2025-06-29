package handlers_test

import (
	"7-solutions/dtos"
	"7-solutions/handlers"
	"7-solutions/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(userDto *dtos.UserRegister) (*dtos.UserResponse, error) {
	args := m.Called(userDto)
	return args.Get(0).(*dtos.UserResponse), args.Error(1)
}

func (m *MockUserService) GetUserByID(id int) (*dtos.UserResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*dtos.UserResponse), args.Error(1)
}

func (m *MockUserService) GetAllUsers() ([]dtos.UserResponse, error) {
	args := m.Called()
	return args.Get(0).([]dtos.UserResponse), args.Error(1)
}

func (m *MockUserService) UpdateUser(id int, userDto *dtos.UserUpdate) (*dtos.UserResponse, error) {
	args := m.Called(id, userDto)
	return args.Get(0).(*dtos.UserResponse), args.Error(1)
}

func (m *MockUserService) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetUserByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(MockUserService)
	h := handlers.NewUserHandler(mockService)
	r.GET("/users/:id", h.GetUserByID)

	user := &models.User{ID: 1, Name: "Test", Email: "test@example.com"}
	mockService.On("GetUserByID", 1).Return(&dtos.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil)

	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Test")
}

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(MockUserService)
	h := handlers.NewUserHandler(mockService)
	r.DELETE("/users/:id", h.DeleteUser)

	mockService.On("DeleteUser", 1).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}

// เพิ่มได้อีกเช่น UpdateUser, GetAllUsers, CreateUser
func TestUpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(MockUserService)
	h := handlers.NewUserHandler(mockService)
	r.PUT("/users/:id", h.UpdateUser)

	userUpdate := &dtos.UserUpdate{Name: "Updated User", Email: "test@example.com"}
	mockService.On("UpdateUser", 1, userUpdate).Return(&dtos.UserResponse{ID: 1, Name: "Updated User", Email: "test@example.com"}, nil)
	payload := `{"name":"Updated User","email":"test@example.com"}`
	req, _ := http.NewRequest(http.MethodPut, "/users/1", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Updated User")
	mockService.AssertCalled(t, "UpdateUser", 1, userUpdate)
}

func TestGetAllUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(MockUserService)
	h := handlers.NewUserHandler(mockService)
	r.GET("/users", h.GetAllUsers)

	users := []dtos.UserResponse{
		{ID: 1, Name: "User1", Email: "test@example.com"},
		{ID: 2, Name: "User2", Email: "test2@example.com"},
	}
	mockService.On("GetAllUsers").Return(users, nil)
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "User1")
	assert.Contains(t, w.Body.String(), "User2")
	mockService.AssertCalled(t, "GetAllUsers")
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(MockUserService)
	h := handlers.NewUserHandler(mockService)
	r.POST("/users", h.CreateUser)

	userDto := &dtos.UserRegister{Name: "New User", Email: "test@example.com", Password: "password123"}
	mockService.On("CreateUser", userDto).Return(&dtos.UserResponse{ID: 1, Name: "New User", Email: "test@example.com"}, nil)
	payload := `{"name":"New User","email":"test@example.com","password":"password123"}`
	req, _ := http.NewRequest(http.MethodPost, "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	assert.Contains(t, w.Body.String(), "New User")
	mockService.AssertCalled(t, "CreateUser", userDto)
}
