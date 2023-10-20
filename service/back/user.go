package back

import (
	"context"
	"time"

	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
	"github.com/lightsaid/hotel-bookings/config"
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
	"github.com/lightsaid/hotel-bookings/pkg/errs"
	"github.com/lightsaid/hotel-bookings/pkg/pswd"
	"github.com/lightsaid/hotel-bookings/pkg/token"
	"github.com/lightsaid/hotel-bookings/service"
)

func (svc *Service) ListUsers(c context.Context, req request.ListRequest) (list []*db.ListUsersRow, total int64, apierr *errs.ApiError) {
	arg := db.ListUsersParams{
		Limit:  req.Limit(),
		Offset: req.Offset(),
	}

	var err error
	list, err = svc.store.ListUsers(c, arg)
	if err != nil {
		apierr = errs.HandleSQLError(err)
		return
	}

	total, _ = svc.store.ListUsersTotal(c)
	return
}

// LoginUser 用户登录，成功返回nil, 错误返回error（错误种类：数据库错误，Token错误, ErrMismatchedPaswd，ErrNotImplemented）
func (svc *Service) LoginUser(c context.Context, req request.LoginRequest) (*reps.LoginResponse, *errs.ApiError) {
	user, err := svc.store.GetUserByPhoneNumber(c, req.PhoneNumber)
	if err != nil {
		return nil, errs.HandleSQLError(err)
	}

	// 验证码登录
	if *req.LoginType == request.LoginType_SMS {
		// TODO:
		return nil, errs.ErrBadRequest.AsMessage(errs.MsgSMSLoginNotImplemented)
	} else {
		// 密码登录
		err = pswd.CheckPassword(req.Password, user.HashedPassword)
		if err != nil {
			return nil, errs.ErrBadRequest.AsMessage(errs.MsgMismatchedPaswd).AsException(err)
		}
	}

	// 验证成功，做登录的其他业务
	aToken, _, err := config.TokenMaker.CreateToken(int64(user.ID), config.Cfg.Token.AccessTokenDuration)
	if err != nil {
		return nil, errs.ErrServerError.AsMessage(errs.MsgCreateTokenFailed).AsException(err)
	}

	rToken, payload, err := config.TokenMaker.CreateToken(int64(user.ID), config.Cfg.Token.RefreshTokenDuration)
	if err != nil {
		return nil, errs.ErrServerError.AsMessage(errs.MsgCreateTokenFailed).AsException(err)
	}

	session := db.CreateSessionParams{
		TokenID:      payload.ID,
		UserID:       user.ID,
		RefreshToken: rToken,
		UserAgent:    req.UserAgent,
		ClientIp:     req.ClientIP,
		ExpiresAt:    payload.ExpiresAt.Time,
		LoginType:    *req.LoginType,
	}
	_, err = service.HandleInsert(c, session, svc.store.CreateSession)
	if err != nil {
		return nil, errs.ErrServerError.AsException(err).AsMessage(errs.MsgCreateSessionFail)
	}

	data := &reps.LoginResponse{
		ID:           user.ID,
		PhoneNumber:  reps.ShadowPhoneNumber(user.PhoneNumber),
		Username:     user.Username,
		Avatar:       user.Avatar,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		AccessToken:  aToken,
		RefreshToken: rToken,
	}

	return data, nil
}

// RenewAccessToken 刷新访问token
func (svc *Service) RenewAccessToken(c context.Context, payload *token.Payload, rToken string) (string, *errs.ApiError) {
	session, err := svc.store.GetSessionByTokenID(c, payload.ID)
	if err := errs.HandleSQLError(err); err != nil {
		return "", err
	}

	if session.RefreshToken != rToken {
		return "", errs.ErrUnauthorized.AsMessage(errs.MsgRefreshTokenNoMis)
	}

	if session.UserID != uint32(payload.UserID) {
		return "", errs.ErrUnauthorized.AsMessage(errs.MsgUserNotMis)
	}

	if time.Now().After(session.ExpiresAt) {
		return "", errs.ErrUnauthorized.AsMessage(token.ErrExpiredToken.Error())
	}

	// 创建access Token
	token, _, err := config.TokenMaker.CreateToken(int64(session.UserID), config.Cfg.Token.AccessTokenDuration)
	if err != nil {
		return "", errs.ErrServerError.AsException(err)
	}

	return token, nil
}
