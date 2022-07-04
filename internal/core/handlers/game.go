package handlers

import "github.com/gofiber/fiber/v2"

type GameHandler interface {
	GameList(c *fiber.Ctx) error
	GameRegis(c *fiber.Ctx) error
	MyGameAccount(c *fiber.Ctx) error
}
