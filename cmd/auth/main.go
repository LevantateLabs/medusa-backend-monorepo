package main

import (
	"context"
	"log"
	"time"

	"github.com/akhil-is-watching/medusa-backend-monorepo/config"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/db"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/jwt"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/logger"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/nats"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/repositories"
	"github.com/akhil-is-watching/medusa-backend-monorepo/service/auth"
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

	authRepository := repositories.NewAuthRepository(db)
	natsClient, err := nats.NewNATSClient(cfg.Nats.Url)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	auth.Run(ctx, cfg, natsClient, authRepository, logger, jwtManager)
}
