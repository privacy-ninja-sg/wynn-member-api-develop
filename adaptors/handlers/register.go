package handlers

import (
	"net/http"
	"wynn-member-api/ent/user"
	"wynn-member-api/internal/core/handlers"
	"wynn-member-api/internal/core/models"
	"wynn-member-api/internal/core/services"
	"wynn-member-api/pkg/errs"
	"wynn-member-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type registerReq struct {
	Tel      string     `json:"tel" validate:"required"`
	Username string     `json:"username" validate:"required"`
	Password string     `json:"password" validate:"required,min=6"`
	Bonus    user.Bonus `json:"bonus" validate:"required"`
	Channel  int        `json:"channel" validate:"required"`
	Token    string     `json:"token" validate:"required"`
	Pin      string     `json:"pin" validate:"required,len=6"`
}

type requestOTPReq struct {
	Tel string `json:"tel" validate:"required,len=10"`
}

type verifyOTPReq struct {
	Token string `json:"token" validate:"required"`
	Pin   string `json:"pin" validate:"required,len=6"`
}

type registerHandler struct {
	registerServe services.RegisterService
}

func NewRegisterHandler(regisService services.RegisterService) handlers.RegisterHandler {
	return &registerHandler{
		registerServe: regisService,
	}
}

func (rh registerHandler) Register(c *fiber.Ctx) error {

	// validate zone
	var reqBody registerReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}
	// wait for otp is coming
	// services zone
	_, err = rh.registerServe.VerifyOTP(reqBody.Pin, reqBody.Token)
	if err != nil {
		return err
	}

	userData, err := rh.registerServe.Register(c.Context(), reqBody.Tel, reqBody.Username, reqBody.Password, reqBody.Bonus, reqBody.Channel)
	if err != nil {
		return errs.NewInternalServerError()
	}

	return c.Status(http.StatusOK).JSON(models.DefaultResponse{
		Status: "ok",
		Code:   http.StatusOK,
		Data:   userData,
	})
}

func (rh registerHandler) RequestOTP(c *fiber.Ctx) error {
	// validate zone
	var reqBody requestOTPReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	resp, err := rh.registerServe.RequestOTP(c.Context(), reqBody.Tel)
	if err != nil {
		return errs.NewBadRequestError(err.Error())
	}

	return c.Status(http.StatusOK).JSON(models.DefaultResponse{
		Status: "ok",
		Code:   http.StatusOK,
		Data: fiber.Map{
			"token": resp.Data.Token,
			"ref":   resp.Data.Ref,
		},
	})
}

func (rh registerHandler) VerifyOTP(c *fiber.Ctx) error {

	// validate zone
	var reqBody verifyOTPReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	resp, err := rh.registerServe.VerifyOTP(reqBody.Pin, reqBody.Token)
	if err != nil {
		return err
	}

	return utils.NewResponse(c, "ok", http.StatusOK, fiber.Map{
		"detail": resp.Data.Detail,
	})
}

func (rh registerHandler) Channels(c *fiber.Ctx) error {
	vals, _ := rh.registerServe.Channels(c.Context())

	return utils.NewResponse(c, "ok", http.StatusOK, vals)
}
