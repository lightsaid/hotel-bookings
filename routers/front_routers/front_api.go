package front_routers

import (
	"github.com/lightsaid/hotel-bookings/api/front"
)

var (
	healthzApi front.HealthZ
	authApi    front.AuthApi
	homeApi    front.HomeApi
)
