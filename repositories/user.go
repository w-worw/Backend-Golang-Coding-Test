package repositories

import (
	"7-solutions/dtos"
	"7-solutions/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUser(userDto *dtos.UserRegister) (*dtos.UserResponse, error)
	GetUserByID(id int) (*dtos.UserResponse, error)
	GetAllUsers() ([]dtos.UserResponse, error)
	UpdateUser(id int, userDto *dtos.UserUpdate) (*dtos.UserResponse, error)
	DeleteUser(id int) error
	CountUsers() (int64, error)
}

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(userDto *dtos.UserRegister) (*dtos.UserResponse, error) {
	newID, err := GetNextSequence(r.db, "users")
	if err != nil {
		return nil, fmt.Errorf("failed to get new user ID: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &models.User{
		ID:        newID,
		Name:      userDto.Name,
		Email:     userDto.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	collection := r.db.Collection("users")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	userResponse := &dtos.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}

	return userResponse, nil
}

func (r *userRepository) GetUserByID(id int) (*dtos.UserResponse, error) {
	var user models.User
	collection := r.db.Collection("users")
	filter := map[string]interface{}{"id": id}
	err := collection.FindOne(nil, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	userResponse := &dtos.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
	return userResponse, nil
}

func (r *userRepository) GetAllUsers() ([]dtos.UserResponse, error) {
	var users []models.User
	collection := r.db.Collection("users")
	cursor, err := collection.Find(nil, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}
	defer cursor.Close(nil)

	for cursor.Next(nil) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("failed to decode user: %w", err)
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	var userResponses []dtos.UserResponse
	for _, user := range users {
		userResponse := dtos.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

func (r *userRepository) UpdateUser(id int, userDto *dtos.UserUpdate) (*dtos.UserResponse, error) {
	user := &models.User{
		ID:    id,
		Name:  userDto.Name,
		Email: userDto.Email,
	}

	collection := r.db.Collection("users")
	filter := map[string]interface{}{"id": id}
	update := map[string]interface{}{
		"$set": user,
	}

	_, err := collection.UpdateOne(nil, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	userResponse := &dtos.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	return userResponse, nil
}

func (r *userRepository) DeleteUser(id int) error {
	collection := r.db.Collection("users")
	filter := map[string]interface{}{"id": id}

	_, err := collection.DeleteOne(nil, filter)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (r *userRepository) CountUsers() (int64, error) {
	collection := r.db.Collection("users")
	count, err := collection.CountDocuments(nil, map[string]interface{}{})
	if err != nil {
		return 0, fmt.Errorf("failed to count users: %w", err)
	}
	return count, nil
}

func GetNextSequence(db *mongo.Database, name string) (int, error) {
	collection := db.Collection("counters")

	filter := bson.M{"_id": name}
	update := bson.M{"$inc": bson.M{"seq": 1}}

	opts := options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After)

	var result struct {
		Seq int `bson:"seq"`
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&result)
	if err != nil {
		return 0, err
	}

	return result.Seq, nil
}
