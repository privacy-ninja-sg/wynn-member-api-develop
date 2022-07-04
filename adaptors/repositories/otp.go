package repositories

import (
	"github.com/parnurzeal/gorequest"
	"wynn-member-api/internal/core/models"
	"wynn-member-api/internal/core/repositories"
)

type otpRepository struct {
	url       string
	apiKey    string
	apiSecret string
}

type ReqOTPBody struct {
	PhoneNumber string `json:"phone_number"`
}

type ReqVerifyOTPBody struct {
	Token string `json:"token"`
	Pin   string `json:"pin"`
}

func NewOTPRepository(url, key, secret string) repositories.OtpRepository {
	return &otpRepository{url, key, secret}
}

func (r otpRepository) Request(tel string) (models.RequestOTPResponse, error) {
	var resp models.RequestOTPResponse

	_, _, errs := gorequest.New().Post(r.url+"/api/request").
		SendStruct(ReqOTPBody{PhoneNumber: tel}).
		SetBasicAuth(r.apiKey, r.apiSecret).
		EndStruct(&resp)

	if errs != nil {
		return models.RequestOTPResponse{}, errs[0]
	}

	return resp, nil
}

func (r otpRepository) Verify(pin, token string) (models.VerifyOTPResponse, error) {
	var resp models.VerifyOTPResponse

	_, _, errs := gorequest.New().Post(r.url+"/api/verify").
		SendStruct(ReqVerifyOTPBody{
			Token: token,
			Pin:   pin,
		}).
		SetBasicAuth(r.apiKey, r.apiSecret).
		EndStruct(&resp)

	if errs != nil {
		return models.VerifyOTPResponse{}, errs[0]
	}

	return resp, nil
}
