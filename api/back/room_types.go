package back

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/api/errs"
	reps "github.com/lightsaid/hotel-bookings/api/response"
)

type RoomTypeApi struct{}

func (*RoomTypeApi) ListRoomTypes(c *gin.Context) {
	list, err := svc.GetListRoomTypes(c)
	if err := errs.HandleSQLError(err); err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, list)
}
