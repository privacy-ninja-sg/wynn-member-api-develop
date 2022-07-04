package repositories

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"wynn-member-api/ent"
	"wynn-member-api/ent/user"
	"wynn-member-api/internal/core/repositories"
)

type userRepository struct {
	db *ent.Client
}

func NewUserRepository(db *ent.Client) repositories.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) Create(ctx context.Context, tel string, username string, password string, bonus user.Bonus, channel int, status user.Status) (*ent.User, error) {
	usr, err := r.db.User.Create().
		SetTel(tel).
		SetUsername(username).
		SetPassword(password).
		SetBonus(bonus).
		SetChannelID(channel).
		SetStatus(status).
		Save(ctx)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return usr, nil
}

func (r userRepository) UpdatePassword(ctx context.Context, userId int, newPassword string) (bool, error) {
	err := r.db.User.UpdateOneID(userId).SetPassword(newPassword).Exec(ctx)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, err
}

func (r userRepository) UpdateBonusStatus(ctx context.Context, userId int, bonus user.Bonus) (bool, error) {
	err := r.db.User.UpdateOneID(userId).SetBonus(bonus).Exec(ctx)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, err
}

func (r userRepository) HasTel(ctx context.Context, tel string) (bool, error) {
	encounter, err := r.db.User.Query().Where(user.Tel(tel)).Count(ctx)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	if encounter > 0 {
		return true, errors.New("account has been already")
	}

	return false, nil
}

func (r userRepository) GetById(ctx context.Context, id int) (*ent.User, error) {
	usr, err := r.db.User.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}

	return usr, nil
}

func (r userRepository) Delete(ctx context.Context, id int) error {
	err := r.db.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r userRepository) UpdateUserStatus(ctx context.Context, userId int, newStatus user.Status) (bool, error) {
	err := r.db.User.UpdateOneID(userId).SetStatus(newStatus).Exec(ctx)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, err
}

func (r userRepository) GetByUsername(ctx context.Context, username string) (*ent.User, error) {
	usr, err := r.db.User.Query().
		Where(user.Username(username)).
		Where(user.StatusEQ(user.StatusActive)).
		First(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return usr, nil
}

//func (u *userRepository) Create(fullName string, tel string, picture string, status user.Status) (*ent.User, error) {
//	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
//	return r.db.User.Create().
//		SetFullName(fullName).
//		SetTel(tel).
//		SetStatus(status).
//		SetPicture(picture).
//		Save(ctx)
//}
//
//func (u *userRepository) UpdateLineClientId(lineClientId string, id int) (bool, error) {
//	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
//	err := r.db.User.UpdateOneID(id).SetLineClientID(lineClientId).Exec(ctx)
//	if err != nil {
//		logrus.Error(err)
//		return false, err
//	}
//	return true, nil
//}
//
//func (u *userRepository) GetByLine(lineClientId string) (*ent.User, error) {
//	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
//	return r.db.User.Query().Where(user.LineClientID(lineClientId)).First(ctx)
//}
//
//func (u *userRepository) GetById(id int) (*ent.User, error) {
//	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
//	return r.db.User.Query().Where(user.ID(id)).Only(ctx)
//
//}
//
//func (u *userRepository) Delete(id int) error {
//	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
//	return r.db.User.DeleteOneID(id).Exec(ctx)
//}
