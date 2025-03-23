package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Patient represents a patient in the system
type Patient struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	AadharNumber string             `bson:"aadharNumber" json:"aadharNumber"`
	Name         string             `bson:"name" json:"name"`
	Age          int                `bson:"age" json:"age"`
	Sex          string             `bson:"sex" json:"sex"`
	Email        string             `bson:"email" json:"email"`
	Phone        string             `bson:"phone" json:"phone"`
	Address      string             `bson:"address" json:"address"`
	MedicalInfo  MedicalInfo        `bson:"medicalInfo" json:"medicalInfo"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// MedicalInfo represents medical information for a patient
type MedicalInfo struct {
	BloodType   string   `bson:"bloodType" json:"bloodType"`
	Allergies   []string `bson:"allergies" json:"allergies"`
	Medications []string `bson:"medications" json:"medications"`
	Conditions  []string `bson:"conditions" json:"conditions"`
}
