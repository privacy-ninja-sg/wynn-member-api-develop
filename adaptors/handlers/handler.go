package handlers

import (
	"github.com/gofiber/fiber/v2"
	"wynn-member-api/internal/core/models"
	"wynn-member-api/pkg/errs"
)

func HandleError(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case errs.AppError:
		return c.Status(e.Code).JSON(models.DefaultResponse{
			Status: "error",
			Code:   errs.ErrorCode(e.Message),
			ErrMsg: e.Message,
		})
	case error:
		return errs.NewInternalServerError()
	default:
		return nil
	}
}
