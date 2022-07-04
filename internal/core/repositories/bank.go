package repositories

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/ent/bank"
)

type BankRepository interface {
	GetBanks(ctx context.Context) ([]*ent.Bank, error)
	GetBank(ctx context.Context, bankId int) (*ent.Bank, error)
	AddBank(ctx context.Context, bankName string, bankShortName string, bankStatus bank.Status, bankLogo string) (*ent.Bank, error)
	DeleteBank(ctx context.Context, bankId int) error
}
