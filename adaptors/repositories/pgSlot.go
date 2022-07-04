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

type pgSlotRepository struct {
	db        *ent.Client
	apiKey    string
	apiSecret string
	host      string
}

func NewPgSlotRepository(db *ent.Client, apiKey string, apiSecret string, host string) repositories.PgSlotRepository {
	return &pgSlotRepository{
		db:        db,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		host:      host,
	}
}

func (r pgSlotRepository) CreateMember(username, password string) (models.PGCreateMemberResponse, error) {
	var resp models.PGCreateMemberResponse
	req := gorequest.New()
	res, _, err := req.Post(r.host + "/api/pgslot/member/create").SendStruct(fiber.Map{
		"username": username,
		"password": password,
	}).EndStruct(&resp)

	if res.StatusCode != http.StatusOK {
		return resp, err[0]
	}

	return resp, nil
}

func (r pgSlotRepository) Deposit(username string, amount float32) (models.PGDepositResponse, error) {
	var resp models.PGDepositResponse
	req := gorequest.New()
	res, _, err := req.Post(r.host + "/api/pgslot/member/deposit").SendStruct(fiber.Map{
		"username": username,
		"amount":   amount,
	}).EndStruct(&resp)

	if res.StatusCode != http.StatusOK {
		return resp, err[0]
	}

	return resp, nil
}

func (r pgSlotRepository) Withdraw(username string, amount float32) (models.PGWithdrawResponse, error) {
	var resp models.PGWithdrawResponse
	req := gorequest.New()
	res, _, err := req.Post(r.host + "/api/pgslot/member/withdraw").SendStruct(fiber.Map{
		"username": username,
		"amount":   amount,
	}).EndStruct(&resp)

	if res.StatusCode != http.StatusOK {
		return resp, err[0]
	}

	return resp, nil
}

func (r pgSlotRepository) GetBalance(username string) (models.PGBalanceResponse, error) {
	var resp models.PGBalanceResponse
	req := gorequest.New()
	res, _, err := req.Post(r.host + "/api/pgslot/member/balance").SendStruct(fiber.Map{
		"username": username,
	}).EndStruct(&resp)

	if res.StatusCode != http.StatusOK {
		return resp, err[0]
	}

	return resp, nil
}

func (r pgSlotRepository) Login(username string) (models.PGLoginResponse, error) {
	var resp models.PGLoginResponse
	req := gorequest.New()
	res, _, err := req.Post(r.host + "/api/pgslot/member/login").SendStruct(fiber.Map{
		"username": username,
	}).EndStruct(&resp)

	if res.StatusCode != http.StatusOK {
		return resp, err[0]
	}

	return resp, nil
}

func (r pgSlotRepository) CreateAccount(ctx context.Context, username, password string, desktopUri, mobileUri string, rawData string, owner int) (*ent.PgSlotAccount, error) {
	return r.db.PgSlotAccount.Create().
		SetUsername(username).
		SetPassword(password).
		SetDesktopURI(desktopUri).
		SetMobileURI(mobileUri).
		SetRawData(rawData).
		SetOwnerID(owner).
		Save(ctx)
}

func (r pgSlotRepository) UpdateGameUrl(ctx context.Context, updateId int, newUrlDesktop, newUrlMobile string) error {
	return r.db.PgSlotAccount.UpdateOneID(updateId).SetDesktopURI(newUrlDesktop).SetMobileURI(newUrlMobile).SetUpdatedAt(time.Now()).Exec(ctx)
}

