package repositories

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/ent/game"
	"wynn-member-api/internal/core/repositories"
)

type gameRepository struct {
	db *ent.Client
}

func NewGameRepository(entCli *ent.Client) repositories.GameRepository {
	return &gameRepository{db: entCli}
}

func (r gameRepository) Get(ctx context.Context) ([]*ent.Game, error) {
	return r.db.Game.Query().Where(game.StatusEQ("on")).All(ctx)
}

func (r gameRepository) GetById(ctx context.Context, gameId int) (*ent.Game, error) {
	return r.db.Game.Get(ctx, gameId)
}
