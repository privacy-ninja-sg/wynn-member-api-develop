package services

import (
	"context"
	"wynn-member-api/ent"
	mwt "wynn-member-api/ent/masterwallettransaction"
)

type WalletService interface {
	GetWalletInfo(ctx context.Context, userId int) (interface{}, error)
	WithdrawBalance(ctx context.Context, userId int, amount float32, remark string) (*ent.MasterWalletTransaction, error)
	RevenueBalance(ctx context.Context, userId int, gameId int) (interface{}, error)
	WalletGameDeposit(ctx context.Context, userId int, gameId int, amount float32) error
	WalletGameWithdraw(ctx context.Context, userId int, gameId int, amount float32) error
	GetWalletTransaction(ctx context.Context, userId int, offset, limit int, txnType mwt.TxnType, status mwt.Status) ([]*ent.MasterWalletTransaction, error)
	GetRevenueAll(ctx context.Context, userId int) (interface{}, error)
}
