package repositories

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
	"wynn-member-api/ent"
	"wynn-member-api/ent/accesstoken"
	"wynn-member-api/ent/user"
	"wynn-member-api/internal/core/repositories"
)

type tokenRepository struct {
	db *ent.Client
}

func NewTokenRepository(db *ent.Client) repositories.TokenRepository {
	return &tokenRepository{db: db}
}

func (r tokenRepository) Create(ctx context.Context, newAccessToken string, tokenExpire time.Time, lineToken string, lineClientID string, uid int, ip string) (*ent.AccessToken, error) {
	exc, err := r.db.AccessToken.Create().
		SetAccessToken(newAccessToken).
		SetLineToken(lineToken).
		SetTokenExpire(tokenExpire).
		SetOwnerID(uid).
		SetIP(ip).
		Save(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return exc, nil
}

func (r tokenRepository) Update(ctx context.Context, newAccessToken string, tokenExpire time.Time, lineToken string, uid int, ip string) (bool, error) {
	_, err := r.db.AccessToken.Update().
		Where(accesstoken.HasOwnerWith(user.ID(uid))).
		SetAccessToken(newAccessToken).
		SetTokenExpire(tokenExpire).
		SetLineToken(lineToken).
		SetUpdatedAt(time.Now()).
		SetIP(ip).
		Save(ctx)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, nil
}

func (r tokenRepository) DeleteByID(ctx context.Context, id int) (bool, error) {
	err := r.db.AccessToken.DeleteOneID(id).Exec(ctx)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	return true, nil
}

func (r tokenRepository) GetByUID(ctx context.Context, uid int) (*ent.AccessToken, error) {
	exc, err := r.db.AccessToken.Query().Where(accesstoken.HasOwnerWith(user.ID(uid))).Only(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return exc, nil
}

func (r tokenRepository) GetByID(ctx context.Context, id int) (*ent.AccessToken, error) {
	exc, err := r.db.AccessToken.Get(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return exc, nil
}
