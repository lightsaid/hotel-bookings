package back

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
	"github.com/lightsaid/hotel-bookings/config"
	"github.com/lightsaid/hotel-bookings/pkg/errs"
	"github.com/lightsaid/hotel-bookings/pkg/token"
)

type AuthApi struct{}

func (*AuthApi) Login(c *gin.Context) {
	var req request.LoginRequest
	if ok := request.ShouldBind(c, &req); !ok {
		return
	}

	req.UserAgent = c.Request.UserAgent()
	req.ClientIP = c.ClientIP()

	data, err := svc.LoginUser(c, req)
	if err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, data)
}

// RenewAccessToken 刷新 AccessToken
func (*AuthApi) RenewAccessToken(c *gin.Context) {
	var req request.RenewAccessTokenRequest
	if ok := request.ShouldBind(c, &req); !ok {
		return
	}

	payload, err := config.TokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		if token.IsCustomError(err) {
			reps.FAIL(c, errs.ErrUnauthorized.AsException(err).AsMessage(err.Error()))
			return
		}
		reps.FAIL(c, errs.ErrUnauthorized.AsException(err))
		return
	}

	token, apierr := svc.RenewAccessToken(c, payload, req.RefreshToken)
	if err != nil {
		reps.FAIL(c, apierr)
		return
	}

	reps.OK(c, token)
}
