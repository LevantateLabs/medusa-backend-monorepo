package response

import "github.com/gofiber/fiber/v2"

func Success(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": data, "error": false})
}

func Error(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error(), "error": true})
}
