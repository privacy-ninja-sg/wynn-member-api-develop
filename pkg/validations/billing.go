package validations

import (
	"wynn-member-api/internal/core/models"
	"wynn-member-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateBillingValidator(c *fiber.Ctx) error {
	var reqBody models.CreateBillingReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	// set request body to locals func
	c.Locals("reqBody", reqBody)

	return c.Next()
}

func UpdateBillingStatusValidator(c *fiber.Ctx) error {
	var reqBody models.UpdateBillingReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	// set request body to locals func
	c.Locals("reqBody", reqBody)

	return c.Next()
}

func GetBillings(c *fiber.Ctx) error {
	var reqBody models.GetBillingReq
	err := c.QueryParser(&reqBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.DefaultResponse{
			Status: "error",
			Code:   http.StatusBadRequest,
			ErrMsg: http.StatusText(http.StatusBadGateway),
		})
	}

	validate := utils.NewValidator()
	err = validate.Struct(reqBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.DefaultResponse{
			Status: "error",
			Code:   http.StatusBadRequest,
			ErrMsg: "validate invalid!",
			Data:   utils.ValidatorErrors(err),
		})
	}

	// set request body to locals func
	c.Locals("reqBody", reqBody)

	return c.Next()
}
