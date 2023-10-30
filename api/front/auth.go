package front

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
	"github.com/lightsaid/hotel-bookings/configs"
	"github.com/lightsaid/hotel-bookings/pkg/errs"
	"github.com/lightsaid/hotel-bookings/pkg/token"
)

type AuthApi struct{}

// Register 注册
func (*AuthApi) Register(c *gin.Context) {
	var req request.ReqisterRequest
	if ok := request.ShouldBind(c, &req); !ok {
		return
	}
	data, err := svc.RegisterUser(c, req)
	if err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, data)
}

// Login 登录
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

	payload, err := configs.TokenMaker.VerifyToken(req.RefreshToken)
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

// GetProfile 获取个人信息
func (*AuthApi) GetProfile(c *gin.Context) {}

// UpdateProfile 更新个人信息
func (*AuthApi) UpdateProfile(c *gin.Context) {}

// BookingRoom 用户订房
func (*AuthApi) BookingRoom(c *gin.Context) {}

// GetBookings 获取用户订房记录
func (*AuthApi) GetBookings(c *gin.Context) {}

// GetBookingsByID 获取个人单个订房记录
func (*AuthApi) GetBookingsByID(c *gin.Context) {}

// GetBookings 用户删除订房记录
func (*AuthApi) DeleteBooking(c *gin.Context) {}

// UserPayment 用户支付
func (*AuthApi) UserPayment(c *gin.Context) {}
