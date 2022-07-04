package repositories

import (
	"context"
	"github.com/sirupsen/logrus"
	"wynn-member-api/ent"
	"wynn-member-api/ent/masterwallettransaction"
	"wynn-member-api/ent/user"
	"wynn-member-api/internal/core/repositories"
)

type masterWalletTransactionRepository struct {
	db *ent.Client
}

func NewMasterWalletTransactionRepository(entCli *ent.Client) repositories.MasterWalletTransactionRepository {
	return &masterWalletTransactionRepository{db: entCli}
}

func (r masterWalletTransactionRepository) GetLatestTransaction(ctx context.Context, userId int) (*ent.MasterWalletTransaction, error) {
	data, err := r.db.MasterWalletTransaction.
		Query().
		Where(masterwallettransaction.HasOwnerWith(user.ID(userId))).
		Order(ent.Desc(masterwallettransaction.FieldID)).
		First(ctx)
	if err != nil {
		logrus.Error(err.Error())
	}
	return data, err
}

func (r masterWalletTransactionRepository) GetLatestBalance(ctx context.Context, userId int) (float32, error) {
	data, err := r.db.MasterWalletTransaction.
		Query().
		Select("balance").
		Where(masterwallettransaction.HasOwnerWith(user.ID(userId))).
		Order(ent.Desc(masterwallettransaction.FieldID)).
		First(ctx)
	if err != nil {
		logrus.Error(err.Error())
	}

	if data == nil {
		return 0, nil
	}
	return data.Balance, nil
}

func (r masterWalletTransactionRepository) CreateTransaction(ctx context.Context, debit, credit, currentBalance float32, txnType masterwallettransaction.TxnType, status masterwallettransaction.Status, userId int, remark string) (*ent.MasterWalletTransaction, error) {
	data, err := r.db.MasterWalletTransaction.Create().
		SetOwnerID(userId).
		SetDebit(debit).
		SetCredit(credit).
		SetBalance(currentBalance).
		SetTxnType(txnType).
		SetStatus(status).
		SetRemark(remark).
		Save(ctx)

	if err != nil {
		logrus.Error(err.Error())
	}

	return data, err
}

func (r masterWalletTransactionRepository) GetTransactionByStatus(ctx context.Context, userId int, status masterwallettransaction.Status) ([]*ent.MasterWalletTransaction, error) {
	data, err := r.db.MasterWalletTransaction.
		Query().
		Where(masterwallettransaction.StatusEQ(status)).
		Where(masterwallettransaction.HasOwnerWith(user.ID(userId))).
		All(ctx)
	if err != nil {
		logrus.Error(err.Error())
	}
	return data, err
}

func (r masterWalletTransactionRepository) UpdateTransaction(ctx context.Context, txnId int, newStatus masterwallettransaction.Status) error {
	_, err := r.db.MasterWalletTransaction.UpdateOneID(txnId).SetStatus(newStatus).Save(ctx)
	return err
}

func (r masterWalletTransactionRepository) GetTransaction(ctx context.Context, userId int, offset, limit int, txnType masterwallettransaction.TxnType, status masterwallettransaction.Status, orderBy ent.OrderFunc) ([]*ent.MasterWalletTransaction, error) {
	return r.db.MasterWalletTransaction.Query().
		Where(masterwallettransaction.HasOwnerWith(
			user.IDEQ(userId),
		)).
		Where(masterwallettransaction.TxnTypeEQ(txnType)).
		Offset(offset).
		Limit(limit).
		Order(orderBy).
		All(ctx)
}
