package repositories

import (
	"context"
	"time"
	"wynn-member-api/ent"
)

type TokenRepository interface {
	Create(ctx context.Context, newAccessToken string, tokenExpire time.Time, lineToken string, lineClientID string, uid int,ip string) (*ent.AccessToken, error)
	Update(ctx context.Context, newAccessToken string, tokenExpire time.Time, lineToken string, uid int, ip string) (bool, error)
	DeleteByID(ctx context.Context, id int) (bool, error)
	GetByUID(ctx context.Context, uid int) (*ent.AccessToken, error)
	GetByID(ctx context.Context, id int) (*ent.AccessToken, error)
}
