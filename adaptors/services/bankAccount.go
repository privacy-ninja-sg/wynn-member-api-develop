package services

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"wynn-member-api/ent"
	"wynn-member-api/ent/bankaccount"
	mwt "wynn-member-api/ent/masterwallettransaction"
	"wynn-member-api/internal/core/repositories"
	"wynn-member-api/internal/core/services"
)

type bankAccountService struct {
	bankAccRepo      repositories.BankAccountRepository
	masterWalletRepo repositories.MasterWalletTransactionRepository
}

func NewBankAccountService(bankAccRepo repositories.BankAccountRepository, masterWalletRepo repositories.MasterWalletTransactionRepository) services.BankAccountService {
	return &bankAccountService{bankAccRepo: bankAccRepo, masterWalletRepo: masterWalletRepo}
}

func (b bankAccountService) CreateBankAccount(ctx context.Context, accName, accId string, userId int, bankCode string) (*ent.BankAccount, error) {
	bankList, _ := b.bankAccRepo.GetByUserId(ctx, userId)
	if len(bankList) != 0 {
		return nil, errors.New("you have already bank account. please contact admin if you need to add more bank account or delete.")
	}

	lastAccId := accId[len(accId)-4:] // last 4 char of string
	return b.bankAccRepo.Create(ctx, accName, accId, lastAccId, bankaccount.StatusApproved, userId, bankCode)
}

func (b bankAccountService) GetBankAccountByUserId(ctx context.Context, userId int) ([]*ent.BankAccount, error) {
	return b.bankAccRepo.GetByUserId(ctx, userId)
}

func (b bankAccountService) DeleteBankAccount(ctx context.Context, userId int, bankId int) (bool, error) {
	return b.bankAccRepo.DeleteBank(ctx, userId, bankId)
}

func (b bankAccountService) CheckBankAccount(ctx context.Context, bankId string) (*ent.BankAccount, error) {
	return b.bankAccRepo.FindBankAccountById(ctx, bankId)
}

func (b bankAccountService) CreditBalance(ctx context.Context, userId int, amount float32, remark string) (*ent.MasterWalletTransaction, error) {
	// deposit function for internal api
	latestBalance, err := b.masterWalletRepo.GetLatestBalance(ctx, userId)
	if err != nil {
		if ent.IsNotFound(err) == false {
			logrus.Error(err.Error())
			return nil, err
		}
	}

	// setup current balance for insert(prepare)
	currentBalance := latestBalance + amount

	data, err := b.masterWalletRepo.CreateTransaction(ctx, amount, 0.0, currentBalance, mwt.TxnTypeDeposit, mwt.StatusSuccessfully, userId, remark)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return data, nil
}

func (b bankAccountService) WithdrawBalance(ctx context.Context, userId int, amount float32, remark string) (*ent.MasterWalletTransaction, error) {
	// withdraw function for member
	panic("implement me")
}
