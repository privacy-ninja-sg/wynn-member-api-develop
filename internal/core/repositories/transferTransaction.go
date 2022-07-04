package repositories

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/ent/transfertransaction"
)

type TransferTransactionRepository interface {
	CreateTransaction(ctx context.Context, userId int, gameId int, amount float32, txnType transfertransaction.TxnType) (*ent.TransferTransaction, error)
	UpdateStatusTransaction(ctx context.Context, txnId int, status transfertransaction.Status) error
	DeleteTransaction(ctx context.Context, txnId int) error
}
