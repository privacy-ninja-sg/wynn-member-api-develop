package handlers

import "github.com/gofiber/fiber/v2"

type TokenHandler interface {
	TokenExchange(c *fiber.Ctx) error
}
