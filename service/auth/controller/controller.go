package controller

import (
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/jwt"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/logger"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/nats"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/repositories"
)

type Controller struct {
	natsClient *nats.NATSClient
	repo       repositories.AuthRepository
	logger     logger.Logger
	jwtManager *jwt.JWTManager
}

func NewBaseController(natsClient *nats.NATSClient, repo repositories.AuthRepository, logger logger.Logger, jwtManager *jwt.JWTManager) *Controller {
	return &Controller{natsClient: natsClient, repo: repo, logger: logger, jwtManager: jwtManager}
}
