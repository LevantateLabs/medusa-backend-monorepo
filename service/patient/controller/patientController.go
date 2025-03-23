package controller

import (
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetPatient(ctx *fiber.Ctx) error {
	aadharNumber := ctx.Locals("aadharNumber").(string)

	patient, err := c.repo.GetPatient(ctx.Context(), aadharNumber)
	if err != nil {
		return response.Error(ctx, err)
	}

	return response.Success(ctx, patient)
}
