package back

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/api/errs"
	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
)

type RoomApi struct{}

func (*RoomApi) CreateRoom(c *gin.Context) {
	var req request.RoomRequest
	if ok := request.ShouldBind(c, &req); !ok {
		return
	}

	newID, err := svc.CreateRoom(c, req)
	if err := errs.HandleSQLError(err); err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, newID)
}

func (*RoomApi) UpdateRoom(c *gin.Context) {
	var req request.RoomRequest
	if ok := request.ShouldBind(c, &req); !ok {
		return
	}

	err := svc.UpdateRoom(c, req)
	if err := errs.HandleSQLError(err); err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, req.ID)
}

func (*RoomApi) ListRooms(c *gin.Context) {
	var req request.ListRequest
	if ok := request.ShouldBind(c, &req, request.BindQuery); !ok {
		return
	}

	list, total, err := svc.ListRooms(c, req)
	if err := errs.HandleSQLError(err); err != nil {
		reps.FAIL(c, err)
		return
	}
	reps.PAGE(c, list, total, req.PageNum, req.PageSize)
}

func (*RoomApi) GetRoom(c *gin.Context) {
	var req request.URIRequest
	if ok := request.ShouldBind(c, &req, request.BindUri); !ok {
		return
	}

	room, err := svc.GetRoomByID(c, req.ID)
	if err := errs.HandleSQLError(err); err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, room)
}

func (*RoomApi) DeleteRoom(c *gin.Context) {
	var req request.URIRequest
	if ok := request.ShouldBind(c, &req, request.BindUri); !ok {
		return
	}
	err := svc.DeleteRoom(c, req.ID)
	if err := errs.HandleSQLError(err); err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, req.ID)
}

func (*RoomApi) UpdateType(c *gin.Context) {
	var req request.UpdateRoomTypeRequest
	if ok := request.ShouldBind(c, &req); !ok {
		return
	}

	err := svc.UpdateRoomType(c, req)
	if err := errs.HandleSQLError(err); err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, req.ID)
}

func (*RoomApi) UpdateStatus(c *gin.Context) {
	var req request.UpdateRoomStatusRequest
	if ok := request.ShouldBind(c, &req); !ok {
		return
	}

	err := svc.UpdateRoomStatus(c, req)
	if err := errs.HandleSQLError(err); err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, req.ID)
}
