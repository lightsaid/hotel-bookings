package front_routers

import (
	"github.com/lightsaid/hotel-bookings/api/front"
	"github.com/lightsaid/hotel-bookings/api/sms"
)

var (
	healthzApi front.HealthZ
	authApi    front.AuthApi
	homeApi    front.HomeApi
	smsApi     sms.SMSApi
)
