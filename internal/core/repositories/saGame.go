package repositories

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/internal/core/models"
)

type SaGameRepository interface {
	// api
	CreateMember(username string) (models.SAGameRegisterResponse, error)
	Login(username string) (models.SAGameLoginResponse, error)
	Profile(username string) (models.SAGameProfileResponse, error)
	Deposit(username string, amount int) (models.SAGameDepositResponse, error)
	Withdraw(username string, amount int) (models.SAGameWithdrawResponse, error)
	// database
	CreateAccount(ctx context.Context, username, password string, desktopUri, mobileUri string, rawData string, owner int) (*ent.SAGameAccount, error)
	UpdateGameUrl(ctx context.Context, updateId int, newUrlDesktop, newUrlMobile string) error

}
