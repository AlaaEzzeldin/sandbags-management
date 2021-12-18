package models


type SendVerifyMailOutput struct {
	OTP string `json:"otp,omitempty"`
}

type Tokens struct {
	RefreshToken string `json:"refresh_token,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	Role         string `json:"role,omitempty"`
}