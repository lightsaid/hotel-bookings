package reps

import "time"

type LoginResponse struct {
	ID           uint32    `json:"id"`
	PhoneNumber  string    `json:"phone_number"`
	Username     string    `json:"username"`
	Avatar       string    `json:"avatar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

// ShadowPhoneNumber 隐藏手机号码中间几位
func ShadowPhoneNumber(phoneNumber string) string {
	if len(phoneNumber) != 11 {
		return phoneNumber
	}

	return phoneNumber[:3] + "*****" + phoneNumber[8:]
}

type RegisterResponse struct {
	ID uint32 `json:"id"`
}
