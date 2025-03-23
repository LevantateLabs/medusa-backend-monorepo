package controller

import (
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetHealth(ctx *fiber.Ctx) error {
	return response.Success(ctx, fiber.Map{
		"message": "OK",
	})
}
