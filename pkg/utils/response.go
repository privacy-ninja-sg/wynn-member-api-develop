package utils

import (
	"github.com/gofiber/fiber/v2"
	"wynn-member-api/internal/core/models"
)

func NewResponse(c *fiber.Ctx, statusText string, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(models.DefaultResponse{
		Status: statusText,
		Code:   statusCode,
		Data:   data,
	})
}
