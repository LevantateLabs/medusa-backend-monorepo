package handler

import (
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/logger"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/nats"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/repositories"
)

type BaseHandler struct {
	Logger      logger.Logger
	authRepo    repositories.AuthRepository
	patientRepo repositories.PatientRepository
	nats        *nats.NATSClient
}

func NewBaseHandler(logger logger.Logger, authRepo repositories.AuthRepository, patientRepo repositories.PatientRepository, nats *nats.NATSClient) *BaseHandler {
	return &BaseHandler{
		Logger:      logger,
		authRepo:    authRepo,
		patientRepo: patientRepo,
		nats:        nats,
	}
}
