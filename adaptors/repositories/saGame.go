package repositories

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
	"wynn-member-api/ent"
	"wynn-member-api/internal/core/models"
	"wynn-member-api/internal/core/repositories"
)

type saGameRepository struct {
	db        *ent.Client
	apiKey    string
	apiSecret string
	host      string
}

func NewSAGameRepository(db *ent.Client, apiKey, apiSecret, host string) repositories.SaGameRepository {
	return &saGameRepository{
		db:        db,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		host:      host,
	}
}

func (r saGameRepository) CreateMember(username string) (models.SAGameRegisterResponse, error) {
	var resp models.SAGameRegisterResponse
	res, _, err := gorequest.New().Post(r.host + "/api/sa/member/create").
		SendStruct(fiber.Map{
			"username": username,
		}).
		EndStruct(&resp)
	if res.StatusCode != http.StatusOK {
		return resp, err[0]
	}

	return resp, nil
}

func (r saGameRepository) Login(username string) (models.SAGameLoginResponse, error) {
	var resp models.SAGameLoginResponse
	res, _, err := gorequest.New().Post(r.host + "/api/sa/member/login").
		SendStruct(fiber.Map{
			"username": username,
		}).
		EndStruct(&resp)
	if res.StatusCode != http.StatusOK {
		return resp, err[0]
	}
	return resp, nil
}

func (r saGameRepository) Profile(username string) (models.SAGameProfileResponse, error) {
	var resp models.SAGameProfileResponse
	res, _, err := gorequest.New().Post(r.host + "/api/sa/member/profile").
		SendStruct(fiber.Map{
			"username": username,
		}).
		EndStruct(&resp)
	if res.StatusCode != http.StatusOK {
		return resp, err[0]
	}
	return resp, nil
}

func (r saGameRepository) Deposit(username string, amount int) (models.SAGameDepositResponse, error) {
	var resp models.SAGameDepositResponse
	res, _, err := gorequest.New().Post(r.host + "/api/sa/member/deposit").
		SendStruct(fiber.Map{
			"username": username,
			"amount":   amount,
		}).
		EndStruct(&resp)
	if res.StatusCode != http.StatusOK {
		return resp, err[0]
	}
	return resp, nil
}

func (r saGameRepository) Withdraw(username string, amount int) (models.SAGameWithdrawResponse, error) {
	var resp models.SAGameWithdrawResponse
	res, _, err := gorequest.New().Post(r.host + "/api/sa/member/withdraw").
		SendStruct(fiber.Map{
			"username": username,
			"amount":   amount,
		}).
		EndStruct(&resp)
	if res.StatusCode != http.StatusOK {
		return resp, err[0]
	}
	return resp, nil
}

func (r saGameRepository) CreateAccount(ctx context.Context, username, password string, desktopUri, mobileUri string, rawData string, owner int) (*ent.SAGameAccount, error) {
	return r.db.SAGameAccount.Create().
		SetUsername(username).
		SetPassword(password).
		SetDesktopURI(desktopUri).
		SetMobileURI(mobileUri).
		SetRawData(rawData).
		SetOwnerID(owner).
		Save(ctx)
}

func (r saGameRepository) UpdateGameUrl(ctx context.Context, updateId int, newUrlDesktop, newUrlMobile string) error {
	return r.db.SAGameAccount.UpdateOneID(updateId).SetDesktopURI(newUrlDesktop).SetMobileURI(newUrlMobile).SetUpdatedAt(time.Now()).Exec(ctx)
}

