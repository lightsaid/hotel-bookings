package back

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/api/errs"
	reps "github.com/lightsaid/hotel-bookings/api/response"
)

type BookingStatusApi struct{}

func (*BookingStatusApi) ListBookingStatus(c *gin.Context) {
	list, err := svc.GetListBookingStatus(c)
	if err := errs.HandleSQLError(err); err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, list)
}
