package handlers

import "github.com/gofiber/fiber/v2"

type BankHandler interface {
	BankList(c *fiber.Ctx) error
	BankCodeList(c *fiber.Ctx) error
}
