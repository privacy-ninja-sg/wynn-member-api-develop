package services

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/internal/core/repositories"
	"wynn-member-api/internal/core/services"
)

type bankService struct {
	bankRepo repositories.BankRepository
}

func NewBankService(bankRepo repositories.BankRepository) services.BankService {
	return &bankService{bankRepo: bankRepo}
}

func (b bankService) GetAllBanks(ctx context.Context) ([]*ent.Bank, error) {
	return b.bankRepo.GetBanks(ctx)
}

func (b bankService) GetBank(ctx context.Context, bankId int) (*ent.Bank, error) {
	return b.bankRepo.GetBank(ctx, bankId)
}
