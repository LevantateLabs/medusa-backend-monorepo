package repositories

import (
	"context"

	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PatientRepository interface {
	CreatePatient(ctx context.Context, patient models.Patient) (models.Patient, error)
	GetPatient(ctx context.Context, aadharNumber string) (models.Patient, error)
	UpdatePatient(ctx context.Context, id string, patient models.Patient) (models.Patient, error)
	DeletePatient(ctx context.Context, id string) error
}

type patientRepository struct {
	db *mongo.Client
}

func NewPatientRepository(db *mongo.Client) PatientRepository {
	return &patientRepository{db: db}
}

func (r *patientRepository) CreatePatient(ctx context.Context, patient models.Patient) (models.Patient, error) {
	collection := r.db.Database("medusa").Collection("patients")
	_, err := collection.InsertOne(ctx, patient)
	return patient, err
}

func (r *patientRepository) GetPatient(ctx context.Context, aadharNumber string) (models.Patient, error) {
	collection := r.db.Database("medusa").Collection("patients")
	var patient models.Patient
	err := collection.FindOne(ctx, bson.M{"aadharNumber": aadharNumber}).Decode(&patient)
	return patient, err
}

func (r *patientRepository) UpdatePatient(ctx context.Context, id string, patient models.Patient) (models.Patient, error) {
	collection := r.db.Database("medusa").Collection("patients")
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": patient})
	return patient, err
}

func (r *patientRepository) DeletePatient(ctx context.Context, id string) error {
	collection := r.db.Database("medusa").Collection("patients")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
