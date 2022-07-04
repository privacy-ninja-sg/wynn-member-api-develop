package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"net/http"
)

func BasicAuth(key, secret string) fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			key: secret,
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == key && pass == secret {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"err": true,
				"msg": http.StatusText(fiber.StatusUnauthorized),
			})
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	})
}
