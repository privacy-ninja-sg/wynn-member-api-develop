package handlers

import "github.com/gofiber/fiber/v2"

type BankAccountHandler interface {
	// public-member-api
	MyBankAccount(c *fiber.Ctx) error
	CreateBankAccount(c *fiber.Ctx) error
	// internal-api
	DeleteBankAccount(c *fiber.Ctx) error
	CheckBankAccount(c *fiber.Ctx) error
	WalletCredit(c *fiber.Ctx) error
}
