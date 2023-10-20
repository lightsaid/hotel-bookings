package front

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
)

type HomeApi struct{}

func (*HomeApi) GetHotels(c *gin.Context) {
	var req request.ListRequest
	if ok := request.ShouldBind(c, &req, request.BindQuery); !ok {
		return
	}
	list, total, err := svc.GetListHotels(c, req)
	if err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.PAGE(c, list, total, req.PageNum, req.PageSize)
}

func (*HomeApi) QueryRooms(c *gin.Context) {
	var req request.QueryRoomsRequest
	if ok := request.ShouldBind(c, &req); !ok {
		return
	}
	list, total, err := svc.QueryRooms(c, req)
	if err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.PAGE(c, list, total, req.PageNum, req.PageSize)
}
