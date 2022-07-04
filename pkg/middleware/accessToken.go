package middleware

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"wynn-member-api/pkg/errs"
	"wynn-member-api/pkg/utils"
)

func AcccessToken(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get now time.
		now := time.Now().Unix()

		// Get claims from JWT.
		claims, err := utils.ExtractTokenMetadata(c, jwtSecret)
		if err != nil {
			// Return status 500 and JWT parse error.
			return errs.NewUnauthorizedError("")
		}

		// Set expiration time from JWT data of current book.
		expires := claims.Expires

		// Checking, if now time greather than expiration from JWT.
		if now > expires {
			// Return status 401 and unauthorized error message.
			return errs.NewUnauthorizedError(errs.EXPIRED_TOKEN)
		}

		// set tokenPayload for handlers
		c.Locals("tokenPayload", claims)

		return c.Next()
	}
}
