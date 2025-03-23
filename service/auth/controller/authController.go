package controller

import (
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"

	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/models"
	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/response"
	"github.com/akhil-is-watching/medusa-backend-monorepo/service/auth/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Controller) CreateUser(ctx *fiber.Ctx) error {
	var req types.CreateAuthRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	model := models.Auth{
		ID:           primitive.NewObjectID(),
		AadharNumber: req.AadharNumber,
		Email:        req.Email,
		Name:         req.Name,
		Phone:        req.Phone,
		Address:      req.Address,
		Age:          req.Age,
		Sex:          req.Sex,
	}

	if _, err := c.repo.CreateAuth(ctx.Context(), model); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	// Convert model to JSON bytes
	modelBytes, err := json.Marshal(model)
	if err != nil {
		return response.Error(ctx, err)
	}

	// Publish the serialized model to NATS
	if err := c.natsClient.Publish("auth.created", modelBytes); err != nil {
		return response.Error(ctx, err)
	}

	return response.Success(ctx, fiber.Map{
		"user": model,
	})
}

func (c *Controller) Authenticate(ctx *fiber.Ctx) error {
	var req types.AuthenticateRequest
	if err := ctx.BodyParser(&req); err != nil {
		c.logger.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	auth, err := c.repo.GetAuthByAadharNumber(ctx.Context(), req.AadharNumber)
	if err != nil {
		c.logger.Error(err)
		return response.Error(ctx, err)
	}

	// Generate a 6-digit OTP between 100000 and 999999
	otp := 100000 + rand.Intn(900000)
	err = c.repo.SetOTP(ctx.Context(), req.AadharNumber, strconv.Itoa(otp))
	if err != nil {
		c.logger.Error(err)
		return response.Error(ctx, err)
	}

	return response.Success(ctx, fiber.Map{
		"auth": auth,
		"otp":  otp,
	})
}

func (c *Controller) VerifyOTP(ctx *fiber.Ctx) error {
	var req types.VerifyOTPRequest
	if err := ctx.BodyParser(&req); err != nil {
		c.logger.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	auth, err := c.repo.GetAuthByAadharNumber(ctx.Context(), req.AadharNumber)
	if err != nil {
		c.logger.Error(err)
		return response.Error(ctx, err)
	}

	if auth.Otp != req.Otp {
		c.logger.Error(errors.New("invalid OTP"))
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid OTP",
		})
	}

	token, err := c.jwtManager.Generate(auth.AadharNumber, auth.Email, "PATIENT")
	if err != nil {
		c.logger.Error(err)
		return response.Error(ctx, err)
	}

	return response.Success(ctx, fiber.Map{
		"auth":  auth,
		"token": token,
	})
}
