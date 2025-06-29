package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        int                `json:"id" bson:"id"`
	ObjectID  primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" validate:"required"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	Password  string             `json:"password" bson:"password" validate:"required,min=6"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}
