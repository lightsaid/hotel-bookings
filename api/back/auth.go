package back

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/api/request"
)

type AuthApi struct{}

func (*AuthApi) Login(c *gin.Context) {
	var req request.LoginRequest
	if ok := request.ShouldBind(c, &req); !ok {
		return
	}

}
