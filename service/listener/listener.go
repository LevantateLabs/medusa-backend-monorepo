package listener

import (
	"context"

	"github.com/akhil-is-watching/medusa-backend-monorepo/config"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/logger"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/nats"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/repositories"
	handler "github.com/akhil-is-watching/medusa-backend-monorepo/service/listener/handlers"
)

func Run(
	ctx context.Context,
	cfg *config.Config,
	natsClient *nats.NATSClient,
	logger logger.Logger,
	authRepo repositories.AuthRepository,
	patientRepo repositories.PatientRepository,
) {

	baseHandler := handler.NewBaseHandler(logger, authRepo, patientRepo, natsClient)
	// Subscribe to the NATS subject with the handler
	_, err := natsClient.Subscribe("auth.created", baseHandler.HandleAuthCreated)
	if err != nil {
		logger.Error(err)
	}

	logger.Info("NATS listener started")

	// Wait for context cancellation (application shutdown)
	<-ctx.Done()
	logger.Info("NATS listener shutting down")
}
