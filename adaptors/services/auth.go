package services

import (
	"context"
	"encoding/hex"
	"errors"
	"github.com/sirupsen/logrus"
	"time"
	"wynn-member-api/ent"
	"wynn-member-api/internal/core/repositories"
	"wynn-member-api/internal/core/services"
	"wynn-member-api/pkg/errs"
	"wynn-member-api/pkg/utils"
)

type authService struct {
	userRepo  repositories.UserRepository
	tokenRepo repositories.TokenRepository
	jwtSecret string
	secretKey []byte
}

func NewAuthService(user repositories.UserRepository, tokenRepo repositories.TokenRepository, jwtSecret string, secret []byte) services.AuthService {
	return &authService{
		userRepo:  user,
		tokenRepo: tokenRepo,
		jwtSecret: jwtSecret,
		secretKey: secret,
	}
}

func (as authService) Login(ctx context.Context, username, password, ip string) (token string, exp int64, err error) {
	userData, err := as.userRepo.GetByUsername(ctx, username)
	if err != nil {
		logrus.Error(err)
		return token, 0, err
	}

	// get object only password for check
	cipherPwdStr := userData.Password

	cipherPwd, _ := hex.DecodeString(cipherPwdStr)
	decryptedPwd, _ := utils.Decrypt(as.secretKey, cipherPwd)

	if string(decryptedPwd) != password {
		return token, 0, errors.New("unauthorized")
	}

	//generate token expire time + 24h
	tokenExp := time.Now().Add(24 * time.Hour)
	// generate access token from content
	token, _ = utils.GenerateNewAccessToken(userData.ID, userData.UUID.String(), userData.Username, "", userData.Status, tokenExp, as.jwtSecret)

	// get token exchange by uid
	exch, err := as.tokenRepo.GetByUID(ctx, userData.ID)
	if err != nil {
		if ent.IsNotFound(err) == false {
			return "", 0, errors.New(errs.GET_TOKEN_DATA)
		}
	}

	if exch != nil {
		// update or create record of exchanges table
		_, err = as.tokenRepo.Update(ctx, token, tokenExp, "", userData.ID, ip)
		if err != nil {
			return "", 0, errors.New(errs.UPDATE_TOKEN_DATA)
		}
	} else {
		// update or create record of exchanges table
		_, err = as.tokenRepo.Create(ctx, token, tokenExp, "", "", userData.ID, ip)
		if err != nil {
			return "", 0, errors.New(errs.CREATE_TOKEN_DATA)
		}
	}

	return token, tokenExp.Unix(), nil
}

func (as authService) ValidatePassword(ctx context.Context, password string, userId int) error {
	userData, err := as.userRepo.GetById(ctx, userId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// get object only password for check
	cipherPwdStr := userData.Password

	cipherPwd, _ := hex.DecodeString(cipherPwdStr)
	decryptedPwd, _ := utils.Decrypt(as.secretKey, cipherPwd)

	if string(decryptedPwd) != password {
		return errors.New("unauthorized")
	}

	return nil
}

func (as authService) EncryptText(text string) ([]byte, error) {
	val, err := utils.Encrypt(as.secretKey, []byte(text))
	if err != nil {
		return []byte{}, err
	}

	return val, err
}
