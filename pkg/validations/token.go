package validations

import (
	"github.com/gofiber/fiber/v2"
	"wynn-member-api/internal/core/models"
	"wynn-member-api/pkg/utils"
)

// TokenExchangeValidator : validator function
func TokenExchangeValidator(c *fiber.Ctx) error {
	var reqBody models.TokenExchangeReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	// set request body to locals func
	c.Locals("reqBody", reqBody)

	return c.Next()
}
