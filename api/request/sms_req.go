package request

type SendRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,vPhone"`
}
