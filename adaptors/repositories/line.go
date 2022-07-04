package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
	"net/http"
	"wynn-member-api/ent"
	"wynn-member-api/internal/core/models"
	"wynn-member-api/internal/core/repositories"
)

type lineRepository struct {
	db       *ent.Client
	clientID string
}

func NewLineRepository(db *ent.Client, clientID string) repositories.LineRepository {
	return &lineRepository{
		db:       db,
		clientID: clientID,
	}
}

func (r lineRepository) Verify(lineIdToken string) (*models.VerifyIDTokenResponse, []error) {
	req := gorequest.New()
	agent := req.Post("https://api.line.me/oauth2/v2.1/verify")
	agent.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	agent.Param("id_token", lineIdToken)
	agent.Param("client_id", r.clientID) // 1656487006
	resp, body, errs := agent.End()
	if errs != nil {
		return nil, errs
	}

	var payload models.VerifyIDTokenResponse

	err := json.Unmarshal([]byte(body), &payload)
	if err != nil {
		logrus.Error(err)
		return nil, []error{err}
	}

	// if error message from line not null.
	if payload.Error != "" {
		logrus.Error(payload.Error, payload.ErrorDescription)
		return nil, []error{errors.New(payload.Error), errors.New(payload.ErrorDescription)}
	}

	if resp.StatusCode != http.StatusOK {
		return nil, []error{errors.New(fmt.Sprintf("http status code error : %d", resp.StatusCode))}
	}

	return &payload, nil
}

func (r lineRepository) UpdateLineClientId(lineClientId string, userId int) (bool, error) {
	panic("implement me")
}

func (r lineRepository) GetUserByLineId(lineClientId string) (*ent.User, error) {
	panic("implement me")
}
