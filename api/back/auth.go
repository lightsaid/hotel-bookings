package back

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
)

type AuthApi struct{}

func (*AuthApi) Login(c *gin.Context) {
	var req request.LoginRequest
	if ok := request.ShouldBind(c, &req); !ok {
		return
	}

	data, err := svc.LoginUser(c, req)
	if err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, data)
}
