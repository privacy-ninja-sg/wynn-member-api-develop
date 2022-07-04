package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"wynn-member-api/internal/core/handlers"
	"wynn-member-api/internal/core/services"
	"wynn-member-api/pkg/errs"
	"wynn-member-api/pkg/utils"
)

type AuthReqBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type authHandler struct {
	authServe services.AuthService
}

func NewAuthHandler(authServe services.AuthService) handlers.AuthHandler {
	return &authHandler{authServe}
}

func (ah authHandler) Login(c *fiber.Ctx) error {
	var reqBody AuthReqBody
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	token, exp, err := ah.authServe.Login(c.Context(), reqBody.Username, reqBody.Password, c.IP())
	if err != nil {
		return errs.NewUnauthorizedError("")
	}

	return utils.NewResponse(c, "ok", http.StatusOK, fiber.Map{"access_token": token, "exp": exp})
}

func (ah authHandler) ForgetPwd(c *fiber.Ctx) error {
	return nil
}
