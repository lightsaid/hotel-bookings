package back

import (
	"github.com/gin-gonic/gin"
	reps "github.com/lightsaid/hotel-bookings/api/response"
)

type BookingStatusApi struct{}

func (*BookingStatusApi) ListBookingStatus(c *gin.Context) {
	list, err := svc.GetListBookingStatus(c)
	if err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, list)
}
