package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Patient represents a patient in the system
type Auth struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	AadharNumber string             `bson:"aadharNumber" json:"aadharNumber"`
	Name         string             `bson:"name" json:"name"`
	Age          int                `bson:"age" json:"age"`
	Sex          string             `bson:"sex" json:"sex"`
	Email        string             `bson:"email" json:"email"`
	Phone        string             `bson:"phone" json:"phone"`
	Address      string             `bson:"address" json:"address"`
	Otp          string             `bson:"otp" json:"otp"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
