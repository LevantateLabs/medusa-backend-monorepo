package repositories

import (
	"context"

	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository interface {
	CreateAuth(ctx context.Context, auth models.Auth) (models.Auth, error)
	GetAuthByAadharNumber(ctx context.Context, aadharNumber string) (models.Auth, error)
	SetOTP(ctx context.Context, aadharNumber string, otp string) error
	UpdateAuth(ctx context.Context, id string, auth models.Auth) (models.Auth, error)
	DeleteAuth(ctx context.Context, id string) error
}

type authRepository struct {
	db *mongo.Client
}

func NewAuthRepository(db *mongo.Client) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateAuth(ctx context.Context, auth models.Auth) (models.Auth, error) {
	collection := r.db.Database("test").Collection("auth")
	_, err := collection.InsertOne(ctx, auth)
	return auth, err
}

func (r *authRepository) GetAuthByAadharNumber(ctx context.Context, aadharNumber string) (models.Auth, error) {
	collection := r.db.Database("test").Collection("auth")
	var auth models.Auth
	err := collection.FindOne(ctx, bson.M{"aadharNumber": aadharNumber}).Decode(&auth)
	return auth, err
}

func (r *authRepository) UpdateAuth(ctx context.Context, id string, auth models.Auth) (models.Auth, error) {
	collection := r.db.Database("test").Collection("auth")
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": auth})
	return auth, err
}

func (r *authRepository) DeleteAuth(ctx context.Context, id string) error {
	collection := r.db.Database("test").Collection("auth")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *authRepository) SetOTP(ctx context.Context, aadharNumber string, otp string) error {
	collection := r.db.Database("test").Collection("auth")
	_, err := collection.UpdateOne(ctx, bson.M{"aadharNumber": aadharNumber}, bson.M{"$set": bson.M{"otp": otp}})
	return err
}
