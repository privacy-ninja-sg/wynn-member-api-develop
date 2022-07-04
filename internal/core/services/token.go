package services

type TokenService interface {
	TokenExchange(idToken string) (accessToken string, exp int64, err error)
}
