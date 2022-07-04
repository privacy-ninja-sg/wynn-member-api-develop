package services

import "context"

type AuthService interface {
	Login(ctx context.Context, username, password, ip string) (token string, exp int64, err error)
	ValidatePassword(ctx context.Context, password string, userId int) error
	EncryptText(text string) ([]byte, error)
}
