package sms

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
	"github.com/lightsaid/hotel-bookings/pkg/smsmock"
)

type SMSApi struct{}

func (*SMSApi) Send(c *gin.Context) {
	sms := smsmock.NewSMS(3 * time.Minute)
	var req request.SendRequest
	if ok := request.ShouldBind(c, &req); !ok {
		return
	}

	smscode, _ := sms.GenSMSCode(req.PhoneNumber)
	reps.OK(c, smscode.Code)
}
