package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"wynn-member-api/ent"
	mwt "wynn-member-api/ent/masterwallettransaction"
	"wynn-member-api/ent/transfertransaction"
	"wynn-member-api/internal/core/repositories"
	"wynn-member-api/internal/core/services"
	"wynn-member-api/pkg/errs"
)

type walletService struct {
	mwtRepo         repositories.MasterWalletTransactionRepository
	prettyGameRepo  repositories.PrettyGameRepository
	pgSlotRepo      repositories.PgSlotRepository
	gameAccRepo     repositories.GameAccountRepository
	transferTxnRepo repositories.TransferTransactionRepository
	saGameRepo      repositories.SaGameRepository
}

func NewWalletService(mwtRepo repositories.MasterWalletTransactionRepository, prettyGameRepo repositories.PrettyGameRepository, pgslotRepo repositories.PgSlotRepository, gameAccRepo repositories.GameAccountRepository, transferTxnRepo repositories.TransferTransactionRepository, saGameRepo repositories.SaGameRepository) services.WalletService {
	return &walletService{
		mwtRepo:         mwtRepo,
		prettyGameRepo:  prettyGameRepo,
		pgSlotRepo:      pgslotRepo,
		gameAccRepo:     gameAccRepo,
		transferTxnRepo: transferTxnRepo,
		saGameRepo:      saGameRepo,
	}
}

func (w walletService) GetWalletInfo(ctx context.Context, userId int) (interface{}, error) {
	var wg sync.WaitGroup

	var reserveBalance float32
	var availableBalance float32

	wg.Add(2)
	go func() {
		defer wg.Done()
		availableBalance, _ = w.mwtRepo.GetLatestBalance(ctx, userId)
	}()

	go func() {
		defer wg.Done()
		txnData, _ := w.mwtRepo.GetTransactionByStatus(ctx, userId, mwt.StatusWaiting)
		for _, txn := range txnData {
			reserveBalance = reserveBalance + txn.Credit
		}
	}()

	wg.Wait()

	return fiber.Map{
		"main_wallet": fiber.Map{
			"available_balance": availableBalance,
			"reserve_balance":   reserveBalance,
		},
	}, nil
}

func (w walletService) WithdrawBalance(ctx context.Context, userId int, amount float32, remark string) (*ent.MasterWalletTransaction, error) {
	// deposit function for internal api
	latestBalance, err := w.mwtRepo.GetLatestBalance(ctx, userId)
	if err != nil {
		if ent.IsNotFound(err) == false {
			logrus.Error(err.Error())
			return nil, err
		}
	}

	if latestBalance < amount {
		// error
		return nil, errors.New("your balance is not enough")
	}

	// setup current balance for insert(prepare)
	currentBalance := latestBalance - amount
	if currentBalance < 0 {
		// error
		return nil, errors.New("invalid balance")
	}

	data, err := w.mwtRepo.CreateTransaction(ctx, 0, amount, currentBalance, mwt.TxnTypeWithdraw, mwt.StatusWaiting, userId, remark)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return data, nil
}

func (w walletService) RevenueBalance(ctx context.Context, userId int, gameId int) (interface{}, error) {
	// find game account from user
	gameAccData, err := w.gameAccRepo.GetGameAccountByUserAndGameID(ctx, userId, gameId)
	if err != nil {
		return nil, errors.New("not found data.")
	}

	switch gameId {
	case 1: // SA
		saData, err := gameAccData.QuerySagame().First(ctx)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		balanceResp, _ := w.saGameRepo.Profile(saData.Username)
		return fiber.Map{
			"balance": balanceResp.Data.Balance,
		}, nil
		break
	case 2: // PG
		pgData, err := gameAccData.QueryPgslot().First(ctx)
		if err != nil {
			return nil, err
		}
		balanceResp, _ := w.pgSlotRepo.GetBalance(pgData.Username)
		return fiber.Map{
			"balance": balanceResp.Data.Balance,
		}, nil
	case 3: // Pretty
		prettyData, err := gameAccData.QueryPretty().First(ctx)
		if err != nil {
			return nil, err
		}
		balanceResp, _ := w.prettyGameRepo.GetBalance(prettyData.Username)
		return fiber.Map{
			"balance": balanceResp.Data.Balance,
		}, nil
	default:
		return nil, errors.New("invalid game id")
	}

	return nil, nil
}

func (w walletService) WalletGameDeposit(ctx context.Context, userId int, gameId int, amount float32) error {
	// get game data
	gameAccData, err := w.gameAccRepo.GetGameAccountByUserAndGameID(ctx, userId, gameId)
	if err != nil {
		return err
	}

	// deposit function for internal api
	latestBalance, err := w.mwtRepo.GetLatestBalance(ctx, userId)
	if err != nil {
		if ent.IsNotFound(err) == false {
			logrus.Error(err.Error())
			return err
		}
	}

	if latestBalance < amount {
		return errors.New("your balance is not enough")
	}

	// deduct balance from main wallet
	currentBalance := latestBalance - amount

	// create transaction to database
	txnMainWallet, err := w.mwtRepo.CreateTransaction(ctx, 0, amount, currentBalance, mwt.TxnTypeTransfer, mwt.StatusPending, userId, fmt.Sprintf("transfer to game : %d", gameId))
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	// create transfer transaction to DB
	transferTxnData, err := w.transferTxnRepo.CreateTransaction(ctx, userId, gameId, amount, transfertransaction.TxnTypeDeposit)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	// deposit amount thb to game
	switch gameId {
	case SAGAME_ID: // SA
		saData, err := gameAccData.QuerySagame().First(ctx)
		if err != nil {
			return err
		}
		resp, err := w.saGameRepo.Deposit(saData.Username, int(amount))
		if err != nil {
			_ = w.mwtRepo.UpdateTransaction(ctx, txnMainWallet.ID, mwt.StatusRejected)
			_ = w.transferTxnRepo.UpdateStatusTransaction(ctx, transferTxnData.ID, transfertransaction.StatusRejected)
			_, _ = w.mwtRepo.CreateTransaction(ctx, amount, 0, currentBalance+amount, mwt.TxnTypeDeposit, mwt.StatusSuccessfully, userId, fmt.Sprintf("refund from game : %d", gameId)) // create this transaction for refund
			return err
		}

		if resp.S != "ok" {
			_ = w.mwtRepo.UpdateTransaction(ctx, txnMainWallet.ID, mwt.StatusRejected)
			_ = w.transferTxnRepo.UpdateStatusTransaction(ctx, transferTxnData.ID, transfertransaction.StatusRejected)
			_, _ = w.mwtRepo.CreateTransaction(ctx, amount, 0, currentBalance+amount, mwt.TxnTypeDeposit, mwt.StatusSuccessfully, userId, fmt.Sprintf("refund from game : %d", gameId)) // create this transaction for refund
			return errors.New(resp.ErrMsg)
		}
		break
	case PGSLOT_ID: // PG
		pgData, err := gameAccData.QueryPgslot().First(ctx)
		if err != nil {
			return err
		}
		resp, err := w.pgSlotRepo.Deposit(pgData.Username, amount)
		if err != nil {
			_ = w.mwtRepo.UpdateTransaction(ctx, txnMainWallet.ID, mwt.StatusRejected)
			_ = w.transferTxnRepo.UpdateStatusTransaction(ctx, transferTxnData.ID, transfertransaction.StatusRejected)
			_, _ = w.mwtRepo.CreateTransaction(ctx, amount, 0, currentBalance+amount, mwt.TxnTypeDeposit, mwt.StatusSuccessfully, userId, fmt.Sprintf("refund from game : %d", gameId)) // create this transaction for refund
			return err
		}
		if resp.S != "ok" {
			_ = w.mwtRepo.UpdateTransaction(ctx, txnMainWallet.ID, mwt.StatusRejected)
			_ = w.transferTxnRepo.UpdateStatusTransaction(ctx, transferTxnData.ID, transfertransaction.StatusRejected)
			_, _ = w.mwtRepo.CreateTransaction(ctx, amount, 0, currentBalance+amount, mwt.TxnTypeDeposit, mwt.StatusSuccessfully, userId, fmt.Sprintf("refund from game : %d", gameId)) // create this transaction for refund
			return errors.New(resp.ErrMsg)
		}
		break
	case PRETTY_ID: // Pretty game
		prettyData, err := gameAccData.QueryPretty().First(ctx)
		if err != nil {
			return err
		}
		resp, err := w.prettyGameRepo.Deposit(prettyData.Username, amount)
		if err != nil {
			_ = w.mwtRepo.UpdateTransaction(ctx, txnMainWallet.ID, mwt.StatusRejected)
			_ = w.transferTxnRepo.UpdateStatusTransaction(ctx, transferTxnData.ID, transfertransaction.StatusRejected)
			_, _ = w.mwtRepo.CreateTransaction(ctx, amount, 0, currentBalance+amount, mwt.TxnTypeDeposit, mwt.StatusSuccessfully, userId, fmt.Sprintf("refund from game : %d", gameId)) // create this transaction for refund
			return err
		}
		if resp.S != "ok" {
			_ = w.mwtRepo.UpdateTransaction(ctx, txnMainWallet.ID, mwt.StatusRejected)
			_ = w.transferTxnRepo.UpdateStatusTransaction(ctx, transferTxnData.ID, transfertransaction.StatusRejected)
			_, _ = w.mwtRepo.CreateTransaction(ctx, amount, 0, currentBalance+amount, mwt.TxnTypeDeposit, mwt.StatusSuccessfully, userId, fmt.Sprintf("refund from game : %d", gameId)) // create this transaction for refund
			return errors.New(resp.ErrMsg)
		}
		break
	default:
		return errors.New("invalid game id")
	}

	// update main wallet transaction status
	_ = w.mwtRepo.UpdateTransaction(ctx, txnMainWallet.ID, mwt.StatusSuccessfully)
	// update transfer transaction table
	_ = w.transferTxnRepo.UpdateStatusTransaction(ctx, transferTxnData.ID, transfertransaction.StatusSuccessfully)

	return nil
}

func (w walletService) WalletGameWithdraw(ctx context.Context, userId int, gameId int, amount float32) error {
	// get game data
	gameAccData, err := w.gameAccRepo.GetGameAccountByUserAndGameID(ctx, userId, gameId)
	if err != nil {
		return err
	}

	var totalBalance float32 = 0
	var pgData *ent.PgSlotAccount
	var ptData *ent.PrettyGameAccount
	var saData *ent.SAGameAccount

	switch gameId {
	case SAGAME_ID: // SA
		saData, err = gameAccData.QuerySagame().First(ctx)
		if err != nil {
			return err
		}

		resp, err := w.saGameRepo.Profile(saData.Username)
		if err != nil {
			return err
		}
		if resp.S != "ok" {
			return errors.New(resp.ErrMsg)
		}

		value, err := strconv.ParseFloat(resp.Data.Balance, 32)
		if err != nil {
			return errs.NewInternalServerError()
		}
		totalBalance = float32(value)
		break
	case PGSLOT_ID: // PG
		pgData, err = gameAccData.QueryPgslot().First(ctx)
		if err != nil {
			return err
		}
		resp, err := w.pgSlotRepo.GetBalance(pgData.Username)
		if err != nil {
			return err
		}
		if resp.S != "ok" {
			return errors.New(resp.ErrMsg)
		}

		totalBalance = resp.Data.Balance
		break
	case PRETTY_ID: // Pretty game
		ptData, err = gameAccData.QueryPretty().First(ctx)
		if err != nil {
			return err
		}
		resp, err := w.prettyGameRepo.GetBalance(ptData.Username)
		if err != nil {
			return err
		}
		if resp.S != "ok" {
			return errors.New(resp.ErrMsg)
		}

		totalBalance = resp.Data.Balance
		break
	default:
		return errors.New("invalid game id")
	}

	if totalBalance < amount {
		return errors.New("your balance is not enough")
	}

	// deposit function for internal api
	latestBalance, err := w.mwtRepo.GetLatestBalance(ctx, userId)
	if err != nil {
		if ent.IsNotFound(err) == false {
			logrus.Error(err.Error())
			return err
		}
	}

	currentBalance := latestBalance + amount

	// create transaction to database
	txnMainWallet, err := w.mwtRepo.CreateTransaction(ctx, amount, 0, currentBalance, mwt.TxnTypeTransfer, mwt.StatusPending, userId, fmt.Sprintf("deposit from game : %d", gameId))
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	// create transfer transaction to DB
	transferTxnData, err := w.transferTxnRepo.CreateTransaction(ctx, userId, gameId, amount, transfertransaction.TxnTypeWithdraw)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	switch gameId {
	case 1: // SA
		resp, err := w.saGameRepo.Withdraw(saData.Username, int(amount))
		if err != nil {
			_ = w.mwtRepo.UpdateTransaction(ctx, txnMainWallet.ID, mwt.StatusRejected)
			_ = w.transferTxnRepo.UpdateStatusTransaction(ctx, transferTxnData.ID, transfertransaction.StatusRejected)
			_, _ = w.mwtRepo.CreateTransaction(ctx, 0, amount, currentBalance-amount, mwt.TxnTypeTransfer, mwt.StatusSuccessfully, userId, fmt.Sprintf("refund to game : %d", gameId)) // create this transaction for refund
			return err
		}
		if resp.S != "ok" {
			_ = w.mwtRepo.UpdateTransaction(ctx, txnMainWallet.ID, mwt.StatusRejected)
			_ = w.transferTxnRepo.UpdateStatusTransaction(ctx, transferTxnData.ID, transfertransaction.StatusRejected)
			_, _ = w.mwtRepo.CreateTransaction(ctx, 0, amount, currentBalance-amount, mwt.TxnTypeTransfer, mwt.StatusSuccessfully, userId, fmt.Sprintf("refund to game : %d", gameId)) // create this transaction for refund
			return errors.New(resp.ErrMsg)
		}
		break
	case 2: // PG
		resp, err := w.pgSlotRepo.Withdraw(pgData.Username, amount)
		if err != nil {
			_ = w.mwtRepo.UpdateTransaction(ctx, txnMainWallet.ID, mwt.StatusRejected)
			_ = w.transferTxnRepo.UpdateStatusTransaction(ctx, transferTxnData.ID, transfertransaction.StatusRejected)
			_, _ = w.mwtRepo.CreateTransaction(ctx, 0, amount, currentBalance-amount, mwt.TxnTypeTransfer, mwt.StatusSuccessfully, userId, fmt.Sprintf("refund to game : %d", gameId)) // create this transaction for refund
			return err
		}
		if resp.S != "ok" {
			_ = w.mwtRepo.UpdateTransaction(ctx, txnMainWallet.ID, mwt.StatusRejected)
			_ = w.transferTxnRepo.UpdateStatusTransaction(ctx, transferTxnData.ID, transfertransaction.StatusRejected)
			_, _ = w.mwtRepo.CreateTransaction(ctx, 0, amount, currentBalance-amount, mwt.TxnTypeTransfer, mwt.StatusSuccessfully, userId, fmt.Sprintf("refund to game : %d", gameId)) // create this transaction for refund
			return errors.New(resp.ErrMsg)
		}

		break
	case 3: // Pretty game
		resp, err := w.prettyGameRepo.Withdraw(ptData.Username, amount)
		if err != nil {
			_ = w.mwtRepo.UpdateTransaction(ctx, txnMainWallet.ID, mwt.StatusRejected)
			_ = w.transferTxnRepo.UpdateStatusTransaction(ctx, transferTxnData.ID, transfertransaction.StatusRejected)
			_, _ = w.mwtRepo.CreateTransaction(ctx, 0, amount, currentBalance-amount, mwt.TxnTypeTransfer, mwt.StatusSuccessfully, userId, fmt.Sprintf("refund to game : %d", gameId)) // create this transaction for refund
			return err
		}
		if resp.S != "ok" {
			_ = w.mwtRepo.UpdateTransaction(ctx, txnMainWallet.ID, mwt.StatusRejected)
			_ = w.transferTxnRepo.UpdateStatusTransaction(ctx, transferTxnData.ID, transfertransaction.StatusRejected)
			_, _ = w.mwtRepo.CreateTransaction(ctx, 0, amount, currentBalance-amount, mwt.TxnTypeTransfer, mwt.StatusSuccessfully, userId, fmt.Sprintf("refund to game : %d", gameId)) // create this transaction for refund
			return errors.New(resp.ErrMsg)
		}

		break
	default:
		return errors.New("invalid game id")
	}

	_ = w.transferTxnRepo.UpdateStatusTransaction(ctx, transferTxnData.ID, transfertransaction.StatusSuccessfully)
	_ = w.mwtRepo.UpdateTransaction(ctx, txnMainWallet.ID, mwt.StatusSuccessfully)

	return nil
}

func (w walletService) GetWalletTransaction(ctx context.Context, userId int, offset, limit int, txnType mwt.TxnType, status mwt.Status) ([]*ent.MasterWalletTransaction, error) {
	return w.mwtRepo.GetTransaction(ctx, userId, offset, limit, txnType, status, ent.Desc(mwt.FieldID))
}

func (w walletService) GetRevenueAll(ctx context.Context, userId int) (interface{}, error) {

	gameAcc, err := w.gameAccRepo.GetGameAccountByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	var balances []fiber.Map

	// init jobs-workers
	var numJobs = len(gameAcc)
	jobs := make(chan *ent.GameAccount, numJobs)
	results := make(chan fiber.Map, numJobs)

	//for _, account := range gameAcc {
	//	switch account.Edges.Game.ID {
	//	case SAGAME_ID: // SA
	//		if len(account.Edges.Sagame) > 0 {
	//			balanceResp, _ := w.saGameRepo.Profile(account.Edges.Sagame[0].Username)
	//			balances = append(balances, fiber.Map{
	//				"game_id":   account.Edges.Game.ID,
	//				"game_name": account.Edges.Game.Name,
	//				"balance":   balanceResp.Data.Balance,
	//			})
	//		}
	//	case PGSLOT_ID: // PG
	//		if len(account.Edges.Pgslot) > 0 {
	//			balanceResp, _ := w.pgSlotRepo.GetBalance(account.Edges.Pgslot[0].Username)
	//			balances = append(balances, fiber.Map{
	//				"game_id":   account.Edges.Game.ID,
	//				"game_name": account.Edges.Game.Name,
	//				"balance":   balanceResp.Data.Balance,
	//			})
	//		}
	//	case PRETTY_ID: // Pretty
	//		if len(account.Edges.Pretty) > 0 {
	//			balanceResp, _ := w.prettyGameRepo.GetBalance(account.Edges.Pretty[0].Username)
	//			balances = append(balances, fiber.Map{
	//				"game_id":   account.Edges.Game.ID,
	//				"game_name": account.Edges.Game.Name,
	//				"balance":   balanceResp.Data.Balance,
	//			})
	//		}
	//	}
	//}

	for wk := 1; wk <= 3; wk++ {
		go worker(wk, jobs, results, w)
	}

	for _, account := range gameAcc {
		jobs <- account
	}
	close(jobs)

	for i := 0; i < numJobs; i++ {
		balances = append(balances, <-results)
	}

	return balances, nil
}

func worker(id int, jobs <-chan *ent.GameAccount, results chan<- fiber.Map, w walletService) {
	for j := range jobs {
		switch j.Edges.Game.ID {
		case SAGAME_ID: // SA
			if len(j.Edges.Sagame) > 0 {
				balanceResp, _ := w.saGameRepo.Profile(j.Edges.Sagame[0].Username)
				results <- fiber.Map{
					"game_id":   j.Edges.Game.ID,
					"game_name": j.Edges.Game.Name,
					"balance":   balanceResp.Data.Balance,
				}
			}
		case PGSLOT_ID: // PG
			if len(j.Edges.Pgslot) > 0 {
				balanceResp, _ := w.pgSlotRepo.GetBalance(j.Edges.Pgslot[0].Username)
				results <- fiber.Map{
					"game_id":   j.Edges.Game.ID,
					"game_name": j.Edges.Game.Name,
					"balance":   balanceResp.Data.Balance,
				}
			}
		case PRETTY_ID: // Pretty
			if len(j.Edges.Pretty) > 0 {
				balanceResp, _ := w.prettyGameRepo.GetBalance(j.Edges.Pretty[0].Username)
				results <- fiber.Map{
					"game_id":   j.Edges.Game.ID,
					"game_name": j.Edges.Game.Name,
					"balance":   balanceResp.Data.Balance,
				}
			}
		}
	}
}
