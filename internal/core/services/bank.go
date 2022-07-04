package services

import (
	"context"
	"wynn-member-api/ent"
)

type BankService interface {
	GetAllBanks(ctx context.Context) ([]*ent.Bank, error)
	GetBank(ctx context.Context, bankId int) (*ent.Bank, error)
}
