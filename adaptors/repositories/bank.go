package repositories

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/ent/bank"
	"wynn-member-api/internal/core/repositories"
)

type bankRepository struct {
	db *ent.Client // postgres
}

func NewBankRepository(entClient *ent.Client) repositories.BankRepository {
	return &bankRepository{db: entClient}
}

func (r bankRepository) GetBanks(ctx context.Context) ([]*ent.Bank, error) {
	return r.db.Bank.Query().Where(bank.StatusEQ(bank.StatusOn)).All(ctx)
}

func (r bankRepository) GetBank(ctx context.Context, bankId int) (*ent.Bank, error) {
	return r.db.Bank.Query().Where(bank.StatusEQ(bank.StatusOn)).Where(bank.ID(bankId)).First(ctx)
}

func (r bankRepository) AddBank(ctx context.Context, bankName string, bankShortName string, bankStatus bank.Status, bankLogo string) (*ent.Bank, error) {
	return r.db.Bank.Create().
		SetName(bankName).
		SetShortName(bankShortName).
		SetStatus(bankStatus).
		SetLogo(bankLogo).
		Save(ctx)
}

func (r bankRepository) DeleteBank(ctx context.Context, bankId int) error {
	return r.db.Bank.DeleteOneID(bankId).Exec(ctx)
}
