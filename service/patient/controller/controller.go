package controller

import (
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/logger"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/repositories"
)

type Controller struct {
	repo   repositories.PatientRepository
	logger logger.Logger
}

func NewBaseController(repo repositories.PatientRepository, logger logger.Logger) *Controller {
	return &Controller{repo: repo, logger: logger}
}
