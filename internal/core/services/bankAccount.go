package services

import (
	"context"
	"wynn-member-api/ent"
)

type BankAccountService interface {
	CreateBankAccount(ctx context.Context, accName, accId string, userId int, bankCode string) (*ent.BankAccount, error)
	GetBankAccountByUserId(ctx context.Context, userId int) ([]*ent.BankAccount, error)
	DeleteBankAccount(ctx context.Context, userId int, bankId int) (bool, error)
	CheckBankAccount(ctx context.Context, bankId string) (*ent.BankAccount, error)
	CreditBalance(ctx context.Context, userId int, amount float32, remark string) (*ent.MasterWalletTransaction, error)   // deposit only
	WithdrawBalance(ctx context.Context, userId int, amount float32, remark string) (*ent.MasterWalletTransaction, error) // withdraw only
}
