package back

import (
	"github.com/gin-gonic/gin"
	reps "github.com/lightsaid/hotel-bookings/api/response"
)

type RoomTypeApi struct{}

func (*RoomTypeApi) ListRoomTypes(c *gin.Context) {
	list, err := svc.GetListRoomTypes(c)
	if err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, list)
}
