package repositories

import (
	"context"
	"wynn-member-api/ent"
)

type GameAccountRepository interface {
	CreateGameAccount(ctx context.Context,userId, gameId int) (*ent.GameAccount, error)
	GetGameAccountByUserAndGameID(ctx context.Context, userId, gameId int) (*ent.GameAccount, error)
	GetGameAccountByUserID(ctx context.Context, userId int) ([]*ent.GameAccount, error)
	DeleteGameAccount(ctx context.Context, userId, gameAccId int) error
}
