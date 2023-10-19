package back

import (
	"context"
	"errors"

	"github.com/lightsaid/hotel-bookings/api/errs"
	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
)

var (
	ErrNotImplemented  = errors.New("验证码登录未实现")
	ErrMismatchedPaswd = errors.New("账号或者密码不匹配")
)

func (svc *Service) ListUsers(c context.Context, req request.ListRequest) (list []*db.ListUsersRow, total int64, err error) {
	arg := db.ListUsersParams{
		Limit:  req.Limit(),
		Offset: req.Offset(),
	}

	list, err = svc.store.ListUsers(c, arg)
	if err != nil {
		return
	}

	total, _ = svc.store.ListUsersTotal(c)
	return
}

// LoginUser 用户登录，成功返回nil, 错误返回error（错误种类：数据库错误，Token错误, ErrMismatchedPaswd，ErrNotImplemented）
func (svc *Service) LoginUser(c context.Context, req request.LoginRequest) (*reps.LoginResponse, *errs.ApiError) {

	// user, err := svc.store.GetUserByPhoneNumber(c, req.PhoneNumber)
	// if err != nil {
	// 	return nil, err
	// }

	// // 验证码登录
	// if req.LoginType == request.LoginType_SMS {
	// 	// TODO:
	// 	return nil, ErrNotImplemented
	// } else {
	// 	// 密码登录
	// 	err = pswd.CheckPassword(req.Password, user.HashedPassword)
	// 	if err != nil {
	// 		return nil, ErrMismatchedPaswd
	// 	}
	// }

	// // 验证成功，做登录的其他业务
	// aToken, _, err := config.TokenMaker.CreateToken(int64(user.ID), config.Cfg.Token.AccessTokenDuration)
	// if err != nil {
	// 	return err
	// }

	return nil, nil
}
