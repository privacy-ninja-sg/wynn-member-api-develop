package services

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/internal/core/models"
)

type GameService interface {
	GetAllGame(ctx context.Context) ([]*ent.Game, error)
	GetGameAccountByGameId(ctx context.Context, userId, gameId int) (*ent.GameAccount, error)
	CreateGameAccount(ctx context.Context, userId, gameId int) (*ent.GameAccount, error)
	GetGameAccount(ctx context.Context, userId int) ([]models.GameAccountListResponse, error)
	DeleteGameAccount(ctx context.Context, userId, gameAccId int) error
}
