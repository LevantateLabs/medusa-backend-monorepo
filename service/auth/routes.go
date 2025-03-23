package auth

import (
	"context"

	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/jwt"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/logger"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/nats"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/repositories"
	"github.com/akhil-is-watching/medusa-backend-monorepo/service/auth/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(_ context.Context, natsClient *nats.NATSClient, authRepository repositories.AuthRepository, logger logger.Logger, jwtManager *jwt.JWTManager) *fiber.App {
	r := fiber.New()

	r.Use(cors.New())

	baseController := controller.NewBaseController(natsClient, authRepository, logger, jwtManager)
	mainRoutes := r.Group("/api")
	{
		mainRoutes.Get("/health", baseController.GetHealth)
		mainRoutes.Post("/auth/create", baseController.CreateUser)
		mainRoutes.Post("/auth/authenticate", baseController.Authenticate)
		mainRoutes.Post("/auth/verify-otp", baseController.VerifyOTP)
	}

	return r
}
