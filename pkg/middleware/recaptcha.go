package middleware

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/ezzarghili/recaptcha-go.v4"
	"time"
	"wynn-member-api/pkg/errs"
	"wynn-member-api/pkg/utils"
)

type SiteVerifyRequest struct {
	RecaptchaResponse string `json:"g-recaptcha-response"`
}

type RecaptchaMiddleware struct {
	SecretKey string `json:"secret_key" validate:"required"`
}

func NewRecaptchaMiddleware(secretKey string) *RecaptchaMiddleware {
	return &RecaptchaMiddleware{SecretKey: secretKey}
}

func (r RecaptchaMiddleware) RecaptchaV2(c *fiber.Ctx) error {
	return c.Next()
	// validate zone
	var reqBody SiteVerifyRequest
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return errs.NewUnauthorizedError("invalid site verify")
	}

	captcha, _ := recaptcha.NewReCAPTCHA(r.SecretKey, recaptcha.V2, 10*time.Second)
	err = captcha.Verify(reqBody.RecaptchaResponse)
	if err != nil {
		return errs.NewUnauthorizedError("please verify your request with recaptcha")
	}
	return c.Next()
}
