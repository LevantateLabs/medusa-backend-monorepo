package main

import (
	"context"
	"log"
	"time"

	"github.com/akhil-is-watching/medusa-backend-monorepo/config"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/db"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/jwt"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/logger"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/repositories"
	"github.com/akhil-is-watching/medusa-backend-monorepo/service/patient"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Failed to load .env file: %v", err)
	// }

	ctx := context.Background()
	cfg := config.LoadConfig()
	logger := logger.NewLogger(cfg.Environment)
	db, err := db.NewMongoClient(cfg.Mongo.Url)
	jwtManager := jwt.NewJWTManager(cfg.JWT.Secret, time.Hour*24)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	patientRepository := repositories.NewPatientRepository(db)
	patient.Run(ctx, cfg, patientRepository, logger, jwtManager)
}
