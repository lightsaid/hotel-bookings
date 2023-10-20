package back_routers

import "github.com/lightsaid/hotel-bookings/api/back"

var (
	healthzApi       back.HealthZ
	authApi          back.AuthApi
	hotelApi         back.HotelApi
	roomApi          back.RoomApi
	roomTypeApi      back.RoomTypeApi
	bookingStatusApi back.BookingStatusApi
	userApi          back.UserApi
	uploadApi        back.UplaodApi
)
