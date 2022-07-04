package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	mwt "wynn-member-api/ent/masterwallettransaction"
	"wynn-member-api/internal/core/handlers"
	"wynn-member-api/internal/core/services"
	"wynn-member-api/pkg/errs"
	"wynn-member-api/pkg/utils"
)

type walletHandler struct {
	walletServe services.WalletService
}

type withdrawReq struct {
	Amount float32 `json:"amount" validate:"required,number,min=100"`
}

type walletRevenueReq struct {
	GameId int `json:"game_id" validate:"required,number"`
}

type walletGameDepositReq struct {
	GameId int     `json:"game_id" validate:"required,number"`
	Amount float32 `json:"amount" validate:"required,number,min=100"`
}

type walletGameWithdrawReq struct {
	GameId int     `json:"game_id" validate:"required,number"`
	Amount float32 `json:"amount" validate:"required,number,min=10"`
}

type historyReq struct {
	Offset string `query:"offset" validate:"required,number,min=0"`
	Limit  string `query:"limit" validate:"required,number,min=1"`
}

func NewWalletHandler(walletServe services.WalletService) handlers.WalletHandler {
	return &walletHandler{walletServe: walletServe}
}

func (w walletHandler) WalletInfo(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)

	walletInfoData, err := w.walletServe.GetWalletInfo(c.Context(), tokenClaims.Uid)
	if err != nil {
		return errs.NewInternalServerError()
	}

	return utils.NewResponse(c, "ok", http.StatusOK, walletInfoData)
}

func (w walletHandler) WalletWithdraw(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)

	// validate zone
	var reqBody withdrawReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	data, err := w.walletServe.WithdrawBalance(c.Context(), tokenClaims.Uid, reqBody.Amount, "NORMAL-WITHDRAW")
	if err != nil {
		return errs.NewBadRequestError(err.Error())
	}

	return utils.NewResponse(c, "ok", http.StatusOK, data)
}

func (w walletHandler) WalletGameDeposit(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)

	// validate zone
	var reqBody walletGameDepositReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	err = w.walletServe.WalletGameDeposit(c.Context(), tokenClaims.Uid, reqBody.GameId, reqBody.Amount)
	if err != nil {
		return errs.NewInternalServerError()
	}

	return utils.NewResponse(c, "ok", http.StatusOK, "successfully")
}

func (w walletHandler) WalletGameWithdraw(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)

	// validate zone
	var reqBody walletGameWithdrawReq
	err := c.BodyParser(&reqBody)
	if err != nil {
		return errs.NewBadRequestError("")
	}

	validate := utils.NewValidator()
	err = validate.Struct(reqBody)
	if err != nil {
		return errs.NewBadRequestError(errs.INVALID_VALIDATION)
	}

	err = w.walletServe.WalletGameWithdraw(c.Context(), tokenClaims.Uid, reqBody.GameId, reqBody.Amount)
	if err != nil {
		return errs.NewInternalServerError()
	}

	return utils.NewResponse(c, "ok", http.StatusOK, "successfully")
}

func (w walletHandler) WalletRevenue(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)

	// validate zone
	var reqBody walletRevenueReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	userId := tokenClaims.Uid
	revenueInfo, err := w.walletServe.RevenueBalance(c.Context(), userId, reqBody.GameId)
	if err != nil {
		return errs.NewNotFoundError(errs.NOT_FOUND)
	}

	return utils.NewResponse(c, "ok", http.StatusOK, revenueInfo)
}

func (w walletHandler) WalletWithdrawHistory(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)

	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return errs.NewBadRequestError(errs.INVALID_VALIDATION)
	}
	limit, err := strconv.Atoi(c.Query("limit", "1"))
	if err != nil {
		return errs.NewBadRequestError(errs.INVALID_VALIDATION)
	}

	results, err := w.walletServe.GetWalletTransaction(c.Context(), tokenClaims.Uid, offset, limit, mwt.TxnTypeWithdraw, mwt.StatusSuccessfully)
	if err != nil {
		return errs.NewInternalServerError()
	}

	return utils.NewResponse(c, "ok", http.StatusOK, results)
}

func (w walletHandler) WalletDepositHistory(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)

	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return errs.NewBadRequestError(errs.INVALID_VALIDATION)
	}
	limit, err := strconv.Atoi(c.Query("limit", "1"))
	if err != nil {
		return errs.NewBadRequestError(errs.INVALID_VALIDATION)
	}

	results, err := w.walletServe.GetWalletTransaction(c.Context(), tokenClaims.Uid, offset, limit, mwt.TxnTypeDeposit, mwt.StatusSuccessfully)
	if err != nil {
		return errs.NewInternalServerError()
	}

	return utils.NewResponse(c, "ok", http.StatusOK, results)
}

func (w walletHandler) WalletTransferHistory(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)
	//
	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return errs.NewBadRequestError(errs.INVALID_VALIDATION)
	}
	limit, err := strconv.Atoi(c.Query("limit", "1"))
	if err != nil {
		return errs.NewBadRequestError(errs.INVALID_VALIDATION)
	}

	results, err := w.walletServe.GetWalletTransaction(c.Context(), tokenClaims.Uid, offset, limit, mwt.TxnTypeTransfer, mwt.StatusSuccessfully)
	if err != nil {
		return errs.NewInternalServerError()
	}

	return utils.NewResponse(c, "ok", http.StatusOK, results)
}

func (w walletHandler) WalletRevenueAll(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)
	allGame, err := w.walletServe.GetRevenueAll(c.Context(), tokenClaims.Uid)
	if err != nil {
		return errs.NewInternalServerError()
	}

	return utils.NewResponse(c, "ok", http.StatusOK, allGame)
}
