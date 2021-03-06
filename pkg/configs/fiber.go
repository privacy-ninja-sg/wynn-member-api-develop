package configs

import (
	"os"
	"strconv"
	"time"
	"wynn-member-api/adaptors/handlers"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Status string      `json:"s"`
	Code   int         `json:"code"`
	Error  interface{} `json:"error"`
}

// FiberConfig func for configuration Fiber adaptors.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ReadTimeout:   time.Second * time.Duration(readTimeoutSecondsCount),
		ErrorHandler:  handlers.HandleError,
		Immutable:     true,
	}
}
