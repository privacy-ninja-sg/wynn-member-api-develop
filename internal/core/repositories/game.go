package repositories

import (
	"context"
	"wynn-member-api/ent"
)

type GameRepository interface {
	Get(ctx context.Context) ([]*ent.Game, error)
	GetById(ctx context.Context, gameId int) (*ent.Game, error)
}
