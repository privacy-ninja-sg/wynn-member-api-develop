package models

type VerifyIDTokenResponse struct {
	Iss     string   `json:"iss,omitempty"`
	Sub     string   `json:"sub,omitempty"`
	Aud     string   `json:"aud,omitempty"`
	Exp     int      `json:"exp,omitempty"`
	Iat     int      `json:"iat,omitempty"`
	Nonce   string   `json:"nonce,omitempty"`
	Amr     []string `json:"amr"` // pwd,linesso,line ,qr
	Name    string   `json:"name,omitempty"`
	Picture string   `json:"picture,omitempty"`
	Email   string   `json:"email,omitempty"`
	// error response
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}
