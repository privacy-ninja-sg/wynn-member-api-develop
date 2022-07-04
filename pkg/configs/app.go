package configs

import "os"

type AppConfig struct {
	AppEnv         string
	AppSecret      string
	OTPUrl         string
	OTPKey         string
	OTPSecret      string
	InternalKey    string
	InternalSecret string
	AgentHost      string
	RecaptchaKey   string
}

func NewAppConfig() AppConfig {
	return AppConfig{
		os.Getenv("APP_ENV"),
		os.Getenv("SECRET_KEY"),
		os.Getenv("OTP_URL"),
		os.Getenv("OTP_API_KEY"),
		os.Getenv("OTP_API_SECRET"),
		os.Getenv("INTERNAL_KEY"),
		os.Getenv("INTERNAL_SECRET"),
		os.Getenv("AGENT_HOST"),
		os.Getenv("RECAPTCHA_SECRET_KEY"),
	}
}
