package models

type RequestOTPResponse struct {
	S    string `json:"s"` // ok, error
	Code int    `json:"code"`
	Data struct {
		Token string `json:"token"`
		Ref   string `json:"ref"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type VerifyOTPResponse struct {
	S    string `json:"s"` // ok, error
	Code int    `json:"code"`
	Data struct {
		Detail string `json:"detail"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}
