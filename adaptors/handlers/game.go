package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"wynn-member-api/ent"
	"wynn-member-api/internal/core/handlers"
	"wynn-member-api/internal/core/services"
	"wynn-member-api/pkg/errs"
	"wynn-member-api/pkg/utils"
)

type gameHandler struct {
	gameServe services.GameService
}

type gameRegisReq struct {
	GameId int `json:"game_id" validate:"required,number"`
}

func NewGameHandler(gameServe services.GameService) handlers.GameHandler {
	return &gameHandler{gameServe: gameServe}
}

func (g gameHandler) GameList(c *fiber.Ctx) error {
	games, err := g.gameServe.GetAllGame(c.Context())
	if err != nil {
		if ent.IsNotFound(err) == false {
			return errs.NewInternalServerError()
		}
	}

	return utils.NewResponse(c, "ok", http.StatusOK, games)
}

func (g gameHandler) GameRegis(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)

	// validate zone
	var reqBody gameRegisReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	gameAcc, err := g.gameServe.GetGameAccountByGameId(c.Context(), tokenClaims.Uid, reqBody.GameId)
	if err != nil {
		if ent.IsNotFound(err) == false {
			return errs.NewNotFoundError(errs.NOT_FOUND)
		}
	}

	if gameAcc != nil {
		return errs.NewBadRequestError(errs.ACCOUNT_ALREADY_CREATED)
	}

	result, err := g.gameServe.CreateGameAccount(c.Context(), tokenClaims.Uid, reqBody.GameId)
	if err != nil {
		return errs.NewInternalServerError()
	}

	return utils.NewResponse(c, "ok", http.StatusOK, result)
}

func (g gameHandler) MyGameAccount(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)

	results, err := g.gameServe.GetGameAccount(c.Context(), tokenClaims.Uid)

	if err != nil {
		return errs.NewInternalServerError()
	}

	return utils.NewResponse(c, "ok", http.StatusOK, results)
}
