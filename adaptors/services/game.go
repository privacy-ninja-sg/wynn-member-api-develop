package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
	"wynn-member-api/ent"
	"wynn-member-api/ent/game"
	"wynn-member-api/internal/core/models"
	"wynn-member-api/internal/core/repositories"
	"wynn-member-api/internal/core/services"
	"wynn-member-api/pkg/utils"
)

type gameService struct {
	gameRepo       repositories.GameRepository
	gameAccRepo    repositories.GameAccountRepository
	prettyGameRepo repositories.PrettyGameRepository
	pgSlotRepo     repositories.PgSlotRepository
	saGameRepo     repositories.SaGameRepository
}

const (
	SAGAME_ID = 1
	PGSLOT_ID = 2
	PRETTY_ID = 3
)

func NewGameService(gameRepo repositories.GameRepository, gameAccRepo repositories.GameAccountRepository, prettyGameRepo repositories.PrettyGameRepository, pgSlotRepo repositories.PgSlotRepository, saGameRepo repositories.SaGameRepository) services.GameService {
	return &gameService{gameRepo: gameRepo, gameAccRepo: gameAccRepo, prettyGameRepo: prettyGameRepo, pgSlotRepo: pgSlotRepo, saGameRepo: saGameRepo}
}

func (g gameService) GetAllGame(ctx context.Context) ([]*ent.Game, error) {
	return g.gameRepo.Get(ctx)
}

func (g gameService) CreateGameAccount(ctx context.Context, userId, gameId int) (*ent.GameAccount, error) {
	// check has game id
	_, err := g.gameRepo.GetById(ctx, gameId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("game id not found")
		}
	}

	// check game account already first
	result, err := g.gameAccRepo.GetGameAccountByUserAndGameID(ctx, userId, gameId)
	if err != nil {
		if ent.IsNotFound(err) == false {
			return nil, err
		}
	}
	if result != nil {
		return nil, errors.New("game account has been already exist")
	}

	gameData, err := g.gameRepo.GetById(ctx, gameId)
	if err != nil {
		return nil, err
	}
	if gameData.Status == game.StatusOff {
		return nil, errors.New("game is closed")
	}

	gameAccData, err := g.gameAccRepo.CreateGameAccount(ctx, userId, gameId)
	if err != nil {
		return nil, err
	}

	// get user data
	userData, err := gameAccData.QueryOwner().First(ctx)
	if err != nil {
		return nil, err
	}

	switch gameId {
	case SAGAME_ID: // SA GAME
		respData, err := g.saGameRepo.CreateMember(userData.Username)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		// get username from data
		username := respData.Data

		rawSaGameData, err := json.Marshal(respData)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		saGameLoginData, err := g.saGameRepo.Login(username)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		_, err = g.saGameRepo.CreateAccount(ctx, username, "", saGameLoginData.Data.Url, saGameLoginData.Data.Url, string(rawSaGameData), gameAccData.ID)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		break
	case PGSLOT_ID: // PG SLOT
		ranPwd := utils.RandString(10)
		pgSlotData, err := g.pgSlotRepo.CreateMember(userData.Username, ranPwd)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		rawPgSlotData, err := json.Marshal(pgSlotData)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		pgSlotLoginData, err := g.pgSlotRepo.Login(pgSlotData.Data.Username)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		_, err = g.pgSlotRepo.CreateAccount(ctx, pgSlotData.Data.Username, ranPwd, pgSlotLoginData.Data.Url, pgSlotLoginData.Data.Url, string(rawPgSlotData), gameAccData.ID)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		break
	case PRETTY_ID: // Pretty GAME
		prettyGameData, err := g.prettyGameRepo.CreateMember(userData.Username, []int{1014, 1015, 1016, 1017, 1018})
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		rawPrettyGameData, err := json.Marshal(prettyGameData)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		_, err = g.prettyGameRepo.CreateAccount(ctx, prettyGameData.Data.PlayerUsername, "", prettyGameData.Data.UriDesktop, prettyGameData.Data.UriMobile, string(rawPrettyGameData), gameAccData.ID)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		break
	}

	return gameAccData, err
}

func (g gameService) GetGameAccount(ctx context.Context, userId int) ([]models.GameAccountListResponse, error) {

	var wrapResposne []models.GameAccountListResponse
	tNow := time.Now()

	var wg sync.WaitGroup

	data, err := g.gameAccRepo.GetGameAccountByUserID(ctx, userId)
	// check login time
	for _, val := range data {
		switch val.Edges.Game.ID {
		case SAGAME_ID: // SA Game
			wg.Add(1)
			go func(gameData *ent.GameAccount, ctx context.Context) {
				defer wg.Done()
				hour := tNow.Sub(gameData.Edges.Sagame[0].UpdatedAt).Minutes() / 60
				if hour > 1 {
					loginData, err := g.saGameRepo.Login(gameData.Edges.Sagame[0].Username) // re-login with username auto
					if err != nil {
						logrus.Error(err)
					}
					err = g.saGameRepo.UpdateGameUrl(ctx, gameData.Edges.Sagame[0].ID, loginData.Data.Url, loginData.Data.Url)
					if err != nil {
						logrus.Error(err)
					}

					// query SA Game
					gameData.Edges.Sagame[0] = val.QuerySagame().FirstX(ctx)

				}
			}(val, ctx)
		case PGSLOT_ID: // PG
			wg.Add(1)
			go func(gameData *ent.GameAccount, ctx context.Context) {
				defer wg.Done()
				hour := tNow.Sub(gameData.Edges.Pgslot[0].UpdatedAt).Minutes() / 60
				if hour > 1 {
					loginData, err := g.pgSlotRepo.Login(gameData.Edges.Pgslot[0].Username)
					if err != nil {
						logrus.Error(err)
					}
					err = g.pgSlotRepo.UpdateGameUrl(ctx, gameData.Edges.Pgslot[0].ID, loginData.Data.Url, loginData.Data.Url)
					if err != nil {
						logrus.Error(err)
					}

					// query PG Slot
					gameData.Edges.Pgslot[0] = val.QueryPgslot().FirstX(ctx)
				}
			}(val, ctx)
		case PRETTY_ID: // Pretty
			wg.Add(1)
			go func(gameData *ent.GameAccount, ctx context.Context) {
				defer wg.Done()
				hour := tNow.Sub(gameData.Edges.Pretty[0].UpdatedAt).Minutes() / 60
				if hour > 1 {
					loginData, err := g.prettyGameRepo.CreateMember(gameData.Edges.Pretty[0].Username, []int{1014, 1015, 1016, 1017, 1018})
					if err != nil {
						logrus.Error(err)
					}
					err = g.prettyGameRepo.UpdateGameUrl(ctx, gameData.Edges.Pretty[0].ID, loginData.Data.UriDesktop, loginData.Data.UriMobile)
					if err != nil {
						logrus.Error(err)
					}

					gameData.Edges.Pretty[0] = val.QueryPretty().FirstX(ctx)
				}
			}(val, ctx)
		}
	}

	wg.Wait()

	// wrap response
	for _, val := range data {
		app := models.GameAccountListResponse{
			Id:        val.ID,
			Uuid:      val.UUID.String(),
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		}

		app.Edges.Game = val.Edges.Game

		switch val.Edges.Game.ID {
		case 1:
			app.Edges.Detail = append(app.Edges.Detail, val.Edges.Sagame[0])
			break
		case 2:
			app.Edges.Detail = append(app.Edges.Detail, val.Edges.Pgslot[0])
			break
		case 3:
			app.Edges.Detail = append(app.Edges.Detail, val.Edges.Pretty[0])
			break
		}
		wrapResposne = append(wrapResposne, app)
	}

	return wrapResposne, err
}

func (g gameService) DeleteGameAccount(ctx context.Context, userId, gameAccId int) error {
	return g.gameAccRepo.DeleteGameAccount(ctx, userId, gameAccId)
}

func (g gameService) GetGameAccountByGameId(ctx context.Context, userId, gameId int) (*ent.GameAccount, error) {
	return g.gameAccRepo.GetGameAccountByUserAndGameID(ctx, userId, gameId)
}
