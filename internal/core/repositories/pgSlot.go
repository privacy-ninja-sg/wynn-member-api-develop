package repositories

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/internal/core/models"
)

type PgSlotRepository interface {
	// api
	CreateMember(username, password string) (models.PGCreateMemberResponse, error)
	Deposit(username string, amount float32) (models.PGDepositResponse, error)
	Withdraw(username string, amount float32) (models.PGWithdrawResponse, error)
	GetBalance(username string) (models.PGBalanceResponse, error)
	Login(username string) (models.PGLoginResponse, error)
	// database
	CreateAccount(ctx context.Context, username, password string, desktopUri, mobileUri string, rawData string, owner int) (*ent.PgSlotAccount, error)
	UpdateGameUrl(ctx context.Context, updateId int, newUrlDesktop, newUrlMobile string) error
}
