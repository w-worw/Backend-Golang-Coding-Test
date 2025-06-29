package repositories

import (
	"7-solutions/dtos"
	"7-solutions/models"
	"7-solutions/utils"

	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository interface {
	RegisterUser(userDto *dtos.UserRegister) error
	AuthenticateUser(input *dtos.UserAuthenticate) (*string, error)
}

type authRepository struct {
	db *mongo.Database
}

func NewAuthRepository(db *mongo.Database) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) RegisterUser(userDto *dtos.UserRegister) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	userDto.Password = string(hashedPassword)
	newID, err := GetNextSequence(r.db, "users")
	if err != nil {
		return fmt.Errorf("failed to get new user ID: %w", err)
	}
	user := &models.User{
		ID:        newID,
		Name:      userDto.Name,
		Email:     userDto.Email,
		Password:  userDto.Password,
		CreatedAt: time.Now(),
	}
	collection := r.db.Collection("users")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *authRepository) AuthenticateUser(input *dtos.UserAuthenticate) (*string, error) {
	var user models.User

	collection := r.db.Collection("users")
	filter := map[string]interface{}{"email": input.Email}
	err := collection.FindOne(nil, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	token, err := utils.GenerateToken(user.Name, user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &token, nil
}
