package request

const (
	LoginType_SMS  = "SMS"
	LoginType_PASS = "PASS"
)

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,vPhone"`
	Password    string `json:"password" binding:"omitempty,min=6,max=16"`
	SmsCode     string `json:"sms_code" binding:"omitempty,min=4,max=8"`
	LoginType   string `json:"login_type" binding:"required,oneof=SMS PASS"`
}
