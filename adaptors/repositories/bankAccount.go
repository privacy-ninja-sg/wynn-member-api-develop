package repositories

import (
	"context"
	"errors"
	"wynn-member-api/ent"
	"wynn-member-api/ent/bankaccount"
	"wynn-member-api/ent/user"
	"wynn-member-api/internal/core/repositories"
)

type bankAccountRepository struct {
	db *ent.Client
}

func NewBankAccountRepository(entClient *ent.Client) repositories.BankAccountRepository {
	return &bankAccountRepository{db: entClient}
}

func (r bankAccountRepository) Create(ctx context.Context, accName, accId, lastAccId string, status bankaccount.Status, userId int, bankCode string) (*ent.BankAccount, error) {
	return r.db.BankAccount.
		Create().
		SetBankAccountID(accId).
		SetBankAccountName(accName).
		SetBankAccountIDLast(lastAccId).
		SetStatus(status).
		SetBankID(1).
		SetBankCode(bankCode).
		SetOwnerID(userId).
		Save(ctx)
}

func (r bankAccountRepository) GetByUserId(ctx context.Context, userId int) ([]*ent.BankAccount, error) {
	return r.db.BankAccount.Query().
		Where(bankaccount.
			HasOwnerWith(user.
				ID(userId),
			),
		).
		WithBank().
		WithBank().
		All(ctx)
}

func (r bankAccountRepository) DeleteBank(ctx context.Context, userId int, bankAccId int) (bool, error) {
	encounter, err := r.db.BankAccount.
		Delete().
		Where(bankaccount.
			HasOwnerWith(user.
				ID(userId),
			),
		).
		Where(bankaccount.ID(bankAccId)).
		Exec(ctx)

	if err != nil {
		return false, err
	}

	if encounter == 0 {
		return false, errors.New("Error : no bank account for delete.")
	}

	return true, nil
}

func (r bankAccountRepository) FindBankAccountById(ctx context.Context, bankId string) (*ent.BankAccount, error) {
	return r.db.BankAccount.Query().Where(bankaccount.BankAccountIDLastEQ(bankId)).WithBank().WithOwner().First(ctx)
}
