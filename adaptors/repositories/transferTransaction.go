package repositories

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/ent/transfertransaction"
	"wynn-member-api/internal/core/repositories"
)

type transferTransactionRepository struct {
	db *ent.Client
}

func NewTransferTransactionRepository(db *ent.Client) repositories.TransferTransactionRepository {
	return &transferTransactionRepository{
		db: db,
	}
}

func (r transferTransactionRepository) CreateTransaction(ctx context.Context, userId int, gameId int, amount float32, txnType transfertransaction.TxnType) (*ent.TransferTransaction, error) {
	return r.db.TransferTransaction.Create().SetOwnerID(userId).SetGameID(gameId).SetAmount(amount).SetStatus(transfertransaction.StatusProcessing).SetTxnType(txnType).Save(ctx)
}

func (r transferTransactionRepository) UpdateStatusTransaction(ctx context.Context, txnId int, status transfertransaction.Status) error {
	_, err := r.db.TransferTransaction.UpdateOneID(txnId).SetStatus(status).Save(ctx)
	return err
}

func (r transferTransactionRepository) DeleteTransaction(ctx context.Context, txnId int) error {
	err := r.db.TransferTransaction.DeleteOneID(txnId).Exec(ctx)
	return err
}
