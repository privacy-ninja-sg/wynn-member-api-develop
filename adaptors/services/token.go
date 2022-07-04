package services

import (
	"wynn-member-api/internal/core/repositories"
	"wynn-member-api/internal/core/services"
)

type tokenService struct {
	line     repositories.LineRepository
	token    repositories.TokenRepository
	user     repositories.UserRepository
	jwtToken string
}

func NewTokenService(line repositories.LineRepository, token repositories.TokenRepository, user repositories.UserRepository, jwtStr string) services.TokenService {
	return &tokenService{
		line:     line,
		token:    token,
		user:     user,
		jwtToken: jwtStr,
	}
}

func (t tokenService) TokenExchange(idToken string) (accessToken string, exp int64, err error) {
	panic("implement me")
}

//func (t *tokenService) TokenExchange(idToken string) (accessToken string, exp int64, err error) {
//	// get line payload from line verify id token
//	linePayload, errs := t.line.Verify(idToken)
//	if errs != nil {
//		return "", 0, errs[1]
//	}
//
//	// get user data from line user id
//	userData, err := t.user.GetByLine(linePayload.Sub)
//
//	if err != nil {
//		if ent.IsNotFound(err) == false {
//			return "", 0, err
//		}
//		// then create user first.
//		userData, err = t.user.Create(linePayload.Name, "", linePayload.Sub, linePayload.Picture, user.StatusInactive)
//		if err != nil {
//			return "", 0, err
//		}
//	}
//
//	// generate token expire time + 24h
//	tokenExp := time.Now().Add(24 * time.Hour)
//	// generate access token from content
//	accToken, _ := utils.GenerateNewAccessToken(userData.ID, userData.UUID.String(), linePayload.Name, linePayload.Picture, userData.Status, tokenExp, t.jwtToken)
//
//	// get token exchange by uid
//	exch, err := t.token.GetByUID(userData.ID)
//	if err != nil {
//		if ent.IsNotFound(err) == false {
//			return "", 0, errors.New(errs.GET_TOKEN_DATA)
//		}
//	}
//	if exch != nil {
//		// update or create record of exchanges table
//		_, err = t.token.Update(accToken, tokenExp, idToken, userData.ID)
//		if err != nil {
//			return "", 0, errors.New(errsUPDATE_TOKEN_DATA)
//		}
//	} else {
//		// update or create record of exchanges table
//		_, err = t.token.Create(accToken, tokenExp, idToken, linePayload.Sub, userData.ID)
//		if err != nil {
//			return "", 0, errors.New(errsCREATE_TOKEN_DATA)
//		}
//	}
//
//	return accToken, tokenExp.Unix(), nil
//}
