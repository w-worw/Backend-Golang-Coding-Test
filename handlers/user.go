package handlers

import (
	"7-solutions/dtos"
	services "7-solutions/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userDto dtos.UserRegister
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	user, err := h.UserService.CreateUser(&userDto)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	c.JSON(201, gin.H{"user": user})
}
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID: " + err.Error()})
		return
	}

	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve users: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"users": users})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID: " + err.Error()})
		return
	}

	var userDto dtos.UserUpdate
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input : " + err.Error()})
		return
	}

	user, err := h.UserService.UpdateUser(id, &userDto)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID: " + err.Error()})
		return
	}

	err = h.UserService.DeleteUser(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user: " + err.Error()})
		return
	}

	c.JSON(204, gin.H{"message": "User deleted successfully"})
}
