package back

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
)

type UserApi struct{}

func (*UserApi) ListUsers(c *gin.Context) {
	var req request.ListRequest
	if ok := request.ShouldBind(c, &req, request.BindQuery); !ok {
		return
	}
	list, total, err := svc.ListUsers(c, req)
	if err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.PAGE(c, list, total, req.PageNum, req.PageSize)
}
