package repositories

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/ent/game"
	"wynn-member-api/ent/gameaccount"
	"wynn-member-api/ent/user"
	"wynn-member-api/internal/core/repositories"
)

type gameAccountRepository struct {
	db *ent.Client
}

func NewGameAccountRepository(entCli *ent.Client) repositories.GameAccountRepository {
	return &gameAccountRepository{db: entCli}
}

func (r gameAccountRepository) CreateGameAccount(ctx context.Context, userId, gameId int) (*ent.GameAccount, error) {
	return r.db.GameAccount.Create().
		SetOwnerID(userId).
		SetGameID(gameId).
		Save(ctx)
}

func (r gameAccountRepository) GetGameAccountByUserAndGameID(ctx context.Context, userId, gameId int) (*ent.GameAccount, error) {
	return r.db.GameAccount.Query().
		Where(gameaccount.
			HasOwnerWith(
				user.ID(userId),
			),
		).
		Where(gameaccount.
			HasGameWith(
				game.ID(gameId),
			),
		).
		First(ctx)
}

func (r gameAccountRepository) GetGameAccountByUserID(ctx context.Context, userId int) ([]*ent.GameAccount, error) {
	return r.db.GameAccount.Query().
		WithGame().
		WithPgslot().
		WithPretty().
		WithSagame().
		Where(gameaccount.
			HasOwnerWith(
				user.ID(userId),
			),
		).All(ctx)
}

func (r gameAccountRepository) DeleteGameAccount(ctx context.Context, userId, gameAccId int) error {
	_, err := r.db.GameAccount.
		Delete().
		Where(
			gameaccount.ID(gameAccId),
		).
		Where(gameaccount.
			HasOwnerWith(
				user.ID(userId),
			),
		).
		Exec(ctx)
	return err
}
