package repositories

import "wynn-member-api/internal/core/models"

type OtpRepository interface {
	Request(tel string) (models.RequestOTPResponse, error)
	Verify(pin, token string) (models.VerifyOTPResponse, error)
}
