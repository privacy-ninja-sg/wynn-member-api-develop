package handlers

import "github.com/gofiber/fiber/v2"

type RegisterHandler interface {
	Register(c *fiber.Ctx) error   // tel., username, password, bonus_status, channel_id
	RequestOTP(c *fiber.Ctx) error // tel.
	VerifyOTP(c *fiber.Ctx) error  // tel., otp, token
	Channels(c *fiber.Ctx) error
}
