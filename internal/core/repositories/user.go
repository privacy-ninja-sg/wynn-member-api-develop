package repositories

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/ent/user"
)

type UserRepository interface {
	Create(ctx context.Context, tel string, username string, password string, bonus user.Bonus, channel int, status user.Status) (*ent.User, error)
	UpdatePassword(ctx context.Context, userId int, newPassword string) (bool, error)
	UpdateBonusStatus(ctx context.Context, userId int, bonus user.Bonus) (bool, error)
	UpdateUserStatus(ctx context.Context, userId int, newStatus user.Status) (bool, error)
	HasTel(ctx context.Context, tel string) (bool, error)
	GetById(ctx context.Context, id int) (*ent.User, error)
	Delete(ctx context.Context, id int) error
	GetByUsername(ctx context.Context, username string) (*ent.User, error)
}
