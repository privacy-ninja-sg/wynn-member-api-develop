package handlers

import "github.com/gofiber/fiber/v2"

type AccountHandler interface {
	Info(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
}
