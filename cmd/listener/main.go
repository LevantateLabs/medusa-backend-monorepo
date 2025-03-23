package main

import (
	"context"
	"log"

	"github.com/akhil-is-watching/medusa-backend-monorepo/config"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/db"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/logger"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/nats"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/repositories"
	"github.com/akhil-is-watching/medusa-backend-monorepo/service/listener"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	cfg := config.LoadConfig()
	logger := logger.NewLogger(cfg.Environment)
	natsClient, err := nats.NewNATSClient(cfg.Nats.Url)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}

	db, err := db.NewMongoClient(cfg.Mongo.Url)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	authRepo := repositories.NewAuthRepository(db)
	patientRepo := repositories.NewPatientRepository(db)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}

	ctx := context.Background()

	listener.Run(ctx, cfg, natsClient, logger, authRepo, patientRepo)
}
