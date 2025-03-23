package patient

import (
	"context"

	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/jwt"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/logger"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/repositories"
	"github.com/akhil-is-watching/medusa-backend-monorepo/service/patient/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(_ context.Context, patientRepository repositories.PatientRepository, logger logger.Logger, jwtManager *jwt.JWTManager) *fiber.App {
	r := fiber.New()

	r.Use(cors.New())

	baseController := controller.NewBaseController(patientRepository, logger)
	mainRoutes := r.Group("/api")
	{
		mainRoutes.Get("/health", baseController.GetHealth)
		mainRoutes.Use(jwtManager.Middleware())
		mainRoutes.Get("/patient", baseController.GetPatient)
	}

	return r
}
