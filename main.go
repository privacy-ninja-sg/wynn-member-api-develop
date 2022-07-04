package main

import (
	"wynn-member-api/ent"
	"wynn-member-api/pkg/configs"
	"wynn-member-api/pkg/errs"
	"wynn-member-api/pkg/routes"
	"wynn-member-api/pkg/utils"
	"wynn-member-api/platform/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	_ = godotenv.Load(".env")
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	// create database connection session & auto-migration
	entCli := database.CreateEntClient()
	defer func(entCli *ent.Client) {
		err := entCli.Close()
		if err != nil {
			logrus.Fatalln("Error! : ", err.Error())
		}
	}(entCli)

	// ========== Initial fiber adaptors ==========
	app := fiber.New(configs.FiberConfig())

	// init router for routing application server...
	appConfig := configs.NewAppConfig()
	app.Use(cors.New())

	router := routes.NewRouter(app, entCli, configs.JwtConfig(), appConfig)
	// run on develop mode
	//router.Swagger()
	router.Monitor()
	router.Logging()
	// run on all mode
	router.Internal() // internal api for admin
	router.Main()     // public-api for member client
	// setup 404
	app.Use(func(c *fiber.Ctx) error {
		return errs.NewNotFoundError(errs.NOT_FOUND)
	})

	// start application server with graceful shutdown...
	utils.StartServerWithGracefulShutdown(app)
}
