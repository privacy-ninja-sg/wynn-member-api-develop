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

type prettyGameRepository struct {
	apiKey    string
	apiSecret string
	host      string
	db        *ent.Client
}

func NewPrettyGameRepository(db *ent.Client, apiKey string, apiSecret string, host string) repositories.PrettyGameRepository {
	return &prettyGameRepository{
		db:        db,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		host:      host,
	}
}

func (r prettyGameRepository) CreateMember(playerUsername string, betLimit []int) (models.PrettyCreateAndLoginResponse, error) {
	var resp models.PrettyCreateAndLoginResponse
	req := gorequest.New()
	res, _, err := req.Post(r.host + "/api/ptg/member/create_login").SendStruct(fiber.Map{
		"playerUsername": playerUsername,
		"betLimit":       betLimit,
	}).EndStruct(&resp)

	if res.StatusCode != http.StatusOK {
		return resp, err[0]
	}

	return resp, nil
}

func (r prettyGameRepository) Deposit(playerUsername string, balance float32) (models.PrettyDepositResponse, error) {
	var resp models.PrettyDepositResponse
	req := gorequest.New()
	res, _, err := req.Post(r.host + "/api/ptg/member/deposit").SendStruct(fiber.Map{
		"playerUsername": playerUsername,
		"balance":        balance,
	}).EndStruct(&resp)

	if res.StatusCode != http.StatusOK {
		return resp, err[0]
	}

	return resp, nil
}

func (r prettyGameRepository) Withdraw(playerUsername string, balance float32) (models.PrettyWithdrawResponse, error) {
	var resp models.PrettyWithdrawResponse
	req := gorequest.New()
	res, _, err := req.Post(r.host + "/api/ptg/member/withdraw").SendStruct(fiber.Map{
		"playerUsername": playerUsername,
		"balance":        balance,
	}).EndStruct(&resp)

	if res.StatusCode != http.StatusOK {
		return resp, err[0]
	}

	return resp, nil
}

func (r prettyGameRepository) GetBalance(playerUsername string) (models.PrettyBalanceResponse, error) {
	var resp models.PrettyBalanceResponse
	req := gorequest.New()
	res, _, err := req.Post(r.host + "/api/ptg/member/balance").SendStruct(fiber.Map{
		"playerUsername": playerUsername,
	}).EndStruct(&resp)

	if res.StatusCode != http.StatusOK {
		return resp, err[0]
	}

	return resp, nil
}

func (r prettyGameRepository) CreateAccount(ctx context.Context, username, password string, desktopUri, mobileUri string, rawData string, owner int) (*ent.PrettyGameAccount, error) {
	return r.db.PrettyGameAccount.Create().
		SetUsername(username).
		SetPassword(password).
		SetDesktopURI(desktopUri).
		SetMobileURI(mobileUri).
		SetRawData(rawData).
		SetOwnerID(owner).
		Save(ctx)
}

func (r prettyGameRepository) UpdateGameUrl(ctx context.Context, updateId int, newUrlDesktop, newUrlMobile string) error {
	return r.db.PrettyGameAccount.UpdateOneID(updateId).SetDesktopURI(newUrlDesktop).SetMobileURI(newUrlMobile).SetUpdatedAt(time.Now()).Exec(ctx)
}
