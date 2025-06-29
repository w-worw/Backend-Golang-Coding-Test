package handlers_test

import (
	"7-solutions/dtos"
	"7-solutions/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) RegisterUser(userDto *dtos.UserRegister) error {
	args := m.Called(userDto)
	return args.Error(0)
}

func (m *MockAuthService) AuthenticateUser(input *dtos.UserAuthenticate) (*string, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*string), args.Error(1)
}

func TestRegisterUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(MockAuthService)
	h := handlers.NewAuthHandler(mockService)

	r.POST("/register", h.RegisterUser)

	user := `{"name":"Test","email":"test@example.com","password":"pass123"}`
	mockService.On("RegisterUser", mock.Anything).Return(nil)

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(user))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	mockService.AssertCalled(t, "RegisterUser", mock.Anything)
}

func TestAuthenticateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(MockAuthService)
	h := handlers.NewAuthHandler(mockService)
	r.POST("/login", h.AuthenticateUser)

	token := "mocked-token"
	mockService.On("AuthenticateUser", mock.Anything).Return(&token, nil)

	payload := `{"email":"test@example.com","password":"pass123"}`
	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "mocked-token")
}
