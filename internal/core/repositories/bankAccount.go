package repositories

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/ent/bankaccount"
)

type BankAccountRepository interface {
	Create(ctx context.Context, accName, accId, lastAccId string, status bankaccount.Status, userId int, bankId string) (*ent.BankAccount, error)
	GetByUserId(ctx context.Context, userId int) ([]*ent.BankAccount, error)
	DeleteBank(ctx context.Context, userId int, bankAccId int) (bool, error)
	FindBankAccountById(ctx context.Context, bankId string) (*ent.BankAccount, error)
}
