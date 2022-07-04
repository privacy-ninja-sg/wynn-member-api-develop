package repositories

import (
	"wynn-member-api/ent"
	"wynn-member-api/internal/core/models"
)

type LineRepository interface {
	Verify(lineIdToken string) (payload *models.VerifyIDTokenResponse, errs []error)
	UpdateLineClientId(lineClientId string, userId int) (bool, error)
	GetUserByLineId(lineClientId string) (*ent.User, error)
}
