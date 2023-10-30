package request

const (
	LoginType_PASS int32 = 0
	LoginType_SMS  int32 = 1
)

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,vPhone"`
	Password    string `json:"password" binding:"omitempty,min=6,max=16"`
	SMSCode     string `json:"sms_code" binding:"omitempty,min=4,max=8"`
	LoginType   *int32 `json:"login_type" binding:"required,oneof=0 1"`
	UserAgent   string `json:"-" binding:"-"`
	ClientIP    string `json:"-" binding:"-"`
}

type RenewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type ReqisterRequest struct {
	UserName    string `json:"username" binding:"required,min=2,max=20"`
	PhoneNumber string `json:"phone_number" binding:"required,vPhone"`
	Password    string `json:"password" binding:"omitempty,min=6,max=16"`
	SMSCode     string `json:"sms_code" binding:"required,min=4,max=10"`
}
