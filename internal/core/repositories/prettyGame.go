package repositories

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/internal/core/models"
)

type PrettyGameRepository interface {
	// api
	CreateMember(playerUsername string, betLimit []int) (models.PrettyCreateAndLoginResponse, error)
	Deposit(playerUsername string, balance float32) (models.PrettyDepositResponse, error)
	Withdraw(playerUsername string, balance float32) (models.PrettyWithdrawResponse, error)
	GetBalance(playerUsername string) (models.PrettyBalanceResponse, error)
	// database
	CreateAccount(ctx context.Context, username, password string, desktopUri, mobileUri string, rawData string, owner int) (*ent.PrettyGameAccount, error)
	UpdateGameUrl(ctx context.Context, updateId int, newUrlDesktop, newUrlMobile string) error
}
