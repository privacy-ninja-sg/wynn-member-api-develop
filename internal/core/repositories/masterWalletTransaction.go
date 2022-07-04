package repositories

import (
	"context"
	"wynn-member-api/ent"
	mwt "wynn-member-api/ent/masterwallettransaction"
)

type MasterWalletTransactionRepository interface {
	GetLatestTransaction(ctx context.Context, userId int) (*ent.MasterWalletTransaction, error)
	GetLatestBalance(ctx context.Context, userId int) (float32, error)
	CreateTransaction(ctx context.Context, debit, credit, currentBalance float32, txnType mwt.TxnType, status mwt.Status, userId int, remark string) (*ent.MasterWalletTransaction, error)
	GetTransactionByStatus(ctx context.Context, userId int, status mwt.Status) ([]*ent.MasterWalletTransaction, error)
	UpdateTransaction(ctx context.Context, txnId int, newStatus mwt.Status) error
	GetTransaction(ctx context.Context, userId int, offset, limit int, txnType mwt.TxnType, status mwt.Status, orderBy ent.OrderFunc) ([]*ent.MasterWalletTransaction, error)
}
