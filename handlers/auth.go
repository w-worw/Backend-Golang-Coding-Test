package handlers

import (
	"7-solutions/dtos"
	services "7-solutions/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var userDto dtos.UserRegister
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	err := h.AuthService.RegisterUser(&userDto)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to register user: " + err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandler) AuthenticateUser(c *gin.Context) {
	var input dtos.UserAuthenticate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	token, err := h.AuthService.AuthenticateUser(&input)
	if err != nil {
		c.JSON(401, gin.H{"error": "Authentication failed: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
