package types

type CreateAuthRequest struct {
	AadharNumber string `json:"aadharNumber"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Age          int    `json:"age"`
	Sex          string `json:"sex"`
}

type AuthenticateRequest struct {
	AadharNumber string `json:"aadharNumber"`
}

type VerifyOTPRequest struct {
	AadharNumber string `json:"aadharNumber"`
	Otp          string `json:"otp"`
}
