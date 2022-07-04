package handlers

import "github.com/gofiber/fiber/v2"

type WalletHandler interface {
	WalletInfo(c *fiber.Ctx) error
	WalletWithdraw(c *fiber.Ctx) error
	WalletGameDeposit(c *fiber.Ctx) error
	WalletGameWithdraw(c *fiber.Ctx) error
	WalletRevenue(c *fiber.Ctx) error
	WalletWithdrawHistory(c *fiber.Ctx) error
	WalletDepositHistory(c *fiber.Ctx) error
	WalletTransferHistory(c *fiber.Ctx) error
	WalletRevenueAll(c *fiber.Ctx) error
}
