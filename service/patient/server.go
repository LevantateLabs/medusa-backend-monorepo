package patient

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/akhil-is-watching/medusa-backend-monorepo/config"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/jwt"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/logger"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/repositories"
)

func Run(ctx context.Context, cfg *config.Config, patientRepository repositories.PatientRepository, logger logger.Logger, jwtManager *jwt.JWTManager) {
	port := fmt.Sprintf(":%d", cfg.Patient.Port)

	router := SetupRoutes(ctx, patientRepository, logger, jwtManager)

	if err := router.Listen(port); err != nil {
		log.Fatalf("Failed to start relayer server: %v", err)
		os.Exit(1)
	}

	log.Printf("listening at 0.0.0.0:%v", port)
}
