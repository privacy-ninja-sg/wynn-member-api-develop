package routes

import (
	"encoding/hex"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"net/http"
	"wynn-member-api/adaptors/handlers"
	"wynn-member-api/adaptors/repositories"
	"wynn-member-api/adaptors/services"
	"wynn-member-api/ent"
	intRepo "wynn-member-api/internal/core/repositories"
	"wynn-member-api/pkg/configs"
	"wynn-member-api/pkg/middleware"
	"wynn-member-api/pkg/utils"
)

type AppRepositories struct {
	tokenRepo        intRepo.TokenRepository
	userRepo         intRepo.UserRepository
	otpRepo          intRepo.OtpRepository
	bankRepo         intRepo.BankRepository
	bankAccRepo      intRepo.BankAccountRepository
	gameRepo         intRepo.GameRepository
	gameAccRepo      intRepo.GameAccountRepository
	masterWalletRepo intRepo.MasterWalletTransactionRepository
	prettyGameRepo   intRepo.PrettyGameRepository
	pgSlotRepo       intRepo.PgSlotRepository
	saGameRepo       intRepo.SaGameRepository
	transferTxnRepo  intRepo.TransferTransactionRepository
	channelRepo      intRepo.ChannelRepository
}

type Router struct {
	app            *fiber.App
	entCli         *ent.Client
	jwtSecret      string
	appEnv         string
	appSecret      []byte
	otpUrl         string
	otpKey         string
	otpSecret      string
	internalKey    string
	internalSecret string
	agentHost      string
	repositories   AppRepositories
	recaptchaKey   string
}

func NewRouter(app *fiber.App, entCli *ent.Client, jwtSecret string, appConfig configs.AppConfig) *Router {

	// new repositories
	appRepositories := AppRepositories{
		tokenRepo:        repositories.NewTokenRepository(entCli),
		userRepo:         repositories.NewUserRepository(entCli),
		otpRepo:          repositories.NewOTPRepository(appConfig.OTPUrl, appConfig.OTPKey, appConfig.OTPSecret),
		bankRepo:         repositories.NewBankRepository(entCli),
		bankAccRepo:      repositories.NewBankAccountRepository(entCli),
		gameRepo:         repositories.NewGameRepository(entCli),
		gameAccRepo:      repositories.NewGameAccountRepository(entCli),
		masterWalletRepo: repositories.NewMasterWalletTransactionRepository(entCli),
		prettyGameRepo:   repositories.NewPrettyGameRepository(entCli, "", "", appConfig.AgentHost),
		pgSlotRepo:       repositories.NewPgSlotRepository(entCli, "", "", appConfig.AgentHost),
		transferTxnRepo:  repositories.NewTransferTransactionRepository(entCli),
		saGameRepo:       repositories.NewSAGameRepository(entCli, "", "", appConfig.AgentHost),
		channelRepo:      repositories.NewChannelRepository(entCli),
	}

	appSecretDec, _ := hex.DecodeString(appConfig.AppSecret)
	return &Router{
		app:            app,
		entCli:         entCli,
		jwtSecret:      jwtSecret,
		appEnv:         appConfig.AppEnv,
		appSecret:      appSecretDec,
		otpUrl:         appConfig.OTPUrl,
		otpKey:         appConfig.OTPKey,
		otpSecret:      appConfig.OTPSecret,
		internalKey:    appConfig.InternalKey,
		internalSecret: appConfig.InternalSecret,
		agentHost:      appConfig.AgentHost,
		repositories:   appRepositories,
		recaptchaKey:   appConfig.RecaptchaKey,
	}
}

// Swagger func for describe group of API Docs routes.
func (r *Router) Swagger() {
	if r.appEnv == "develop" {
		// Create routes group.
		route := r.app.Group("/swagger")
		// Route for GET method:
		route.Get("*", middleware.BasicAuth(r.internalKey, r.internalSecret), swagger.Handler) // get one user by ID
	}
}

func (r *Router) Monitor() {
	if r.appEnv == "develop" {
		r.app.Get("/monitor", middleware.BasicAuth(r.internalKey, r.internalSecret), monitor.New())
	}
}

func (r *Router) Logging() {
	if r.appEnv == "develop" {
		r.app.Use(logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}))
	}
}

func (r *Router) Internal() {
	// services
	bankAccServe := services.NewBankAccountService(r.repositories.bankAccRepo, r.repositories.masterWalletRepo)
	// handlers
	bankAccHandler := handlers.NewBankAccountHandler(bankAccServe)

	// api group
	api := r.app.Group("api")
	wynnInternal := api.Group("wynn/internal")

	// basic auth for internal api
	basicAuthMiddleware := middleware.BasicAuth(r.internalKey, r.internalSecret)

	// wynn-internal
	wynnInternal.Use(basicAuthMiddleware)

	wynnInternal.Post("wallet/credit", bankAccHandler.WalletCredit)
	wynnInternal.Post("bank/check", bankAccHandler.CheckBankAccount)
	wynnInternal.Post("bank/account/delete", bankAccHandler.DeleteBankAccount) // delete bank account from admin
}

func (r *Router) Main() {

	repo := r.repositories

	// new services
	authServe := services.NewAuthService(repo.userRepo, repo.tokenRepo, r.jwtSecret, r.appSecret)
	registerServe := services.NewRegisterService(repo.otpRepo, repo.userRepo, repo.channelRepo, r.appSecret)
	bankServe := services.NewBankService(repo.bankRepo)
	bankAccServe := services.NewBankAccountService(repo.bankAccRepo, repo.masterWalletRepo)
	gameServe := services.NewGameService(repo.gameRepo, repo.gameAccRepo, repo.prettyGameRepo, repo.pgSlotRepo, repo.saGameRepo)
	walletServe := services.NewWalletService(repo.masterWalletRepo, repo.prettyGameRepo, repo.pgSlotRepo, repo.gameAccRepo, repo.transferTxnRepo, repo.saGameRepo)

	// new handlers
	authHandler := handlers.NewAuthHandler(authServe)
	registerHandler := handlers.NewRegisterHandler(registerServe)
	bankHandler := handlers.NewBankHandler(bankServe)
	bankAccHandler := handlers.NewBankAccountHandler(bankAccServe)
	gameHandler := handlers.NewGameHandler(gameServe)
	walletHandler := handlers.NewWalletHandler(walletServe)
	accountHandler := handlers.NewAccountHandler(r.entCli, authServe, repo.userRepo)

	// api group
	api := r.app.Group("api")
	wynnAuth := api.Group("wynn")

	// recaptcha
	//rcpt := middleware.NewRecaptchaMiddleware(r.recaptchaKey)

	// un-auth endpoint
	// register
	api.Post("register/otp/request", registerHandler.RequestOTP) // with recaptcha
	api.Post("register/otp/verify", registerHandler.VerifyOTP)
	api.Post("register/create", registerHandler.Register) // with recaptcha
	api.Get("register/channels", registerHandler.Channels)

	// auth-login
	api.Post("auth/login", authHandler.Login) // with recaptcha

	// require-auth-zone /account
	authMiddleware := middleware.AcccessToken(r.jwtSecret)

	// wynn-auth
	wynnAuth.Use(authMiddleware) // add auth middleware to recheck

	wynnAuth.Get("bank/list", bankHandler.BankList) // get wynn bank for deposit auto
	wynnAuth.Get("bank_code/list", bankHandler.BankCodeList)
	wynnAuth.Get("account/info", accountHandler.Info)
	wynnAuth.Post("account/change-pwd", accountHandler.ChangePassword)

	// wallet
	wynnAuth.Get("wallet/info", walletHandler.WalletInfo)                               // get wallet info about balance
	wynnAuth.Post("wallet/withdraw", walletHandler.WalletWithdraw)                      // withdraw THB to bank
	wynnAuth.Post("wallet/game/deposit", walletHandler.WalletGameDeposit)               // transfer amount to game
	wynnAuth.Post("wallet/game/withdraw", walletHandler.WalletGameWithdraw)             // transfer amount from game to main wallet
	wynnAuth.Post("wallet/game/revenue", walletHandler.WalletRevenue)                   // get revenue by games
	wynnAuth.Get("wallet/game/revenue/all", walletHandler.WalletRevenueAll)             // get revenue all game
	wynnAuth.Get("wallet/deposit/histories", walletHandler.WalletDepositHistory)        // get deposit history by offset, limit
	wynnAuth.Get("wallet/withdraw/histories", walletHandler.WalletWithdrawHistory)      // get withdraw history by offset, limit
	wynnAuth.Get("wallet/game/transfer/histories", walletHandler.WalletTransferHistory) // get transfer history by offset, limit

	// bank
	wynnAuth.Get("bank/account/list", bankAccHandler.MyBankAccount)        // get available bank account
	wynnAuth.Post("bank/account/create", bankAccHandler.CreateBankAccount) // set available bank account

	// game
	wynnAuth.Get("game/list", gameHandler.GameList) // get available game list
	wynnAuth.Get("game/account/list", gameHandler.MyGameAccount)
	wynnAuth.Post("game/regis", gameHandler.GameRegis) // regis game for user owner

	r.app.Get("", func(c *fiber.Ctx) error {
		return utils.NewResponse(c, "ok", http.StatusOK, fiber.Map{
			"s":   "ok",
			"msg": "hello world",
		})
	})
}
