package services

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/ent/user"
	"wynn-member-api/internal/core/models"
)

type RegisterService interface {
	Register(ctx context.Context, tel string, username string, password string, bonus user.Bonus, channel int) (*ent.User, error)
	RequestOTP(ctx context.Context, tel string) (models.RequestOTPResponse, error)
	VerifyOTP(pin, token string) (models.VerifyOTPResponse, error)
	Channels(ctx context.Context) ([]*ent.Channel, error)
}
