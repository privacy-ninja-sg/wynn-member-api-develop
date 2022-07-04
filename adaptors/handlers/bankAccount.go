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

type bankAccountHandler struct {
	bankAccServe services.BankAccountService
}

type createBankReq struct {
	AccName  string `json:"acc_name" validate:"required"`
	AccId    string `json:"acc_id" validate:"required,max=20"`
	BankCode string `json:"bank_code" validate:"required,min=3,max=3"`
}

type deleteBankReq struct {
	UserId    int `json:"user_id" validate:"required,number"`
	BankAccId int `json:"bank_acc_id" validate:"required,number"`
}

type checkBankReq struct {
	LastBankId string `json:"last_bank_id" validate:"required,max=4"`
}

type depositWalletReq struct {
	WalletId int     `json:"wallet_id" validate:"required,number"`
	Amount   float32 `json:"amount" validate:"required,number"`
	Remark   string  `json:"remark"`
}

func NewBankAccountHandler(bankAccServe services.BankAccountService) handlers.BankAccountHandler {
	return &bankAccountHandler{
		bankAccServe: bankAccServe,
	}
}

func (b bankAccountHandler) MyBankAccount(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)

	bankAccList, err := b.bankAccServe.GetBankAccountByUserId(c.Context(), tokenClaims.Uid)
	if err != nil {
		if ent.IsNotFound(err) == false {
			return errs.NewNotFoundError(errs.NOT_FOUND)
		}
	}
	return utils.NewResponse(c, "ok", http.StatusOK, bankAccList)
}

func (b bankAccountHandler) CreateBankAccount(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)

	// validate zone
	var reqBody createBankReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	result, err := b.bankAccServe.CreateBankAccount(c.Context(), reqBody.AccName, reqBody.AccId, tokenClaims.Uid, reqBody.BankCode)
	if err != nil {
		return errs.NewInternalServerError()
	}

	return utils.NewResponse(c, "ok", http.StatusOK, result)
}

func (b bankAccountHandler) DeleteBankAccount(c *fiber.Ctx) error {
	// validate zone
	var reqBody deleteBankReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	_, err = b.bankAccServe.DeleteBankAccount(c.Context(), reqBody.UserId, reqBody.BankAccId)
	if err != nil {
		return errs.NewInternalServerError()
	}

	return utils.NewResponse(c, "ok", http.StatusOK, nil)
}

func (b bankAccountHandler) CheckBankAccount(c *fiber.Ctx) error {
	// validate zone
	var reqBody checkBankReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	bankAcc, _ := b.bankAccServe.CheckBankAccount(c.Context(), reqBody.LastBankId)

	return utils.NewResponse(c, "ok", http.StatusOK, bankAcc)
}

func (b bankAccountHandler) WalletCredit(c *fiber.Ctx) error {
	// validate zone
	var reqBody depositWalletReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	// credit services
	data, err := b.bankAccServe.CreditBalance(c.Context(), reqBody.WalletId, reqBody.Amount, reqBody.Remark)
	if err != nil {
		return errs.NewInternalServerError()
	}

	return utils.NewResponse(c, "ok", http.StatusOK, data)
}
