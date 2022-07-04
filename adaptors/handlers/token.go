package handlers

import (
	"github.com/gofiber/fiber/v2"
	"wynn-member-api/internal/core/handlers"
	"wynn-member-api/internal/core/services"
)

type tokenHandler struct {
	ts services.TokenService
}

func NewTokenHandler(tokenService services.TokenService) handlers.TokenHandler {
	return &tokenHandler{
		ts: tokenService,
	}
}

func (t tokenHandler) TokenExchange(c *fiber.Ctx) error {
	panic("implement me")
}

//func (th *tokenHandler) TokenExchange(c *fiber.Ctx) error {
//	reqBodyStr := c.Locals("reqBody")
//	reqBody := reqBodyStr.(models.TokenExchangeReq)
//
//	accessToken, exp, err := th.ts.TokenExchange(reqBody.IdToken)
//	if err != nil {
//		return c.Status(http.StatusInternalServerError).JSON(models.DefaultResponse{
//			Status: "error",
//			Code:   http.StatusInternalServerError,
//			ErrMsg: err.Error(),
//		})
//	}
//
//	return c.Status(http.StatusOK).JSON(models.DefaultResponse{
//		Status: "ok",
//		Code:   http.StatusOK,
//		Data: models.TokenExchangeResponse{
//			AccessToken: accessToken,
//			Expire:      exp,
//		},
//	})
//}
