package controller

import (
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetHealth(ctx *fiber.Ctx) error {
	c.natsClient.Publish("auth.created", []byte("test"))
	return response.Success(ctx, fiber.Map{
		"message": "OK",
	})
}
