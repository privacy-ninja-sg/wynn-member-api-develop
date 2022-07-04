package services

import (
	"context"
	"encoding/hex"
	"errors"
	"wynn-member-api/ent"
	"wynn-member-api/ent/user"
	"wynn-member-api/internal/core/models"
	"wynn-member-api/internal/core/repositories"
	"wynn-member-api/internal/core/services"
	"wynn-member-api/pkg/utils"
)

type registerService struct {
	otp       repositories.OtpRepository
	user      repositories.UserRepository
	chann     repositories.ChannelRepository
	secretKey []byte
}

func NewRegisterService(otp repositories.OtpRepository, user repositories.UserRepository, chann repositories.ChannelRepository, secretKey []byte) services.RegisterService {
	return &registerService{
		otp:       otp,
		user:      user,
		chann:     chann,
		secretKey: secretKey,
	}
}

func (rs registerService) Register(ctx context.Context, tel string, username string, password string, bonus user.Bonus, channel int) (*ent.User, error) {
	// encrypted password
	encryptPwd, err := utils.Encrypt(rs.secretKey, []byte(password))
	cipherPwd := hex.EncodeToString(encryptPwd)

	resp, err := rs.user.Create(ctx, tel, username, cipherPwd, bonus, channel, user.StatusActive)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (rs registerService) RequestOTP(ctx context.Context, tel string) (models.RequestOTPResponse, error) {
	hasStatus, err := rs.user.HasTel(ctx, tel)
	if hasStatus {
		return models.RequestOTPResponse{}, errors.New("account has been already")
	}

	if err != nil {
		if ent.IsNotFound(err) == false {
			return models.RequestOTPResponse{}, err
		}
	}

	resp, err := rs.otp.Request(tel)
	if err != nil {
		return models.RequestOTPResponse{}, err
	}

	return resp, nil
}

func (rs registerService) VerifyOTP(pin, token string) (models.VerifyOTPResponse, error) {
	resp, err := rs.otp.Verify(pin, token)
	if err != nil {
		return models.VerifyOTPResponse{}, err
	}

	return resp, nil
}

func (rs registerService) Channels(ctx context.Context) ([]*ent.Channel, error) {
	return rs.chann.GetChannels(ctx)
}
