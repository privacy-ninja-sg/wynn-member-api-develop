package models

type TokenExchangeReq struct {
	IdToken string `json:"id_token" validate:"required"`
}
