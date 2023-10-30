package front

import (
	"context"
	"time"

	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
	"github.com/lightsaid/hotel-bookings/configs"
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
	"github.com/lightsaid/hotel-bookings/pkg/errs"
	"github.com/lightsaid/hotel-bookings/pkg/pswd"
	"github.com/lightsaid/hotel-bookings/pkg/smsmock"
	"github.com/lightsaid/hotel-bookings/pkg/token"
	"github.com/lightsaid/hotel-bookings/service"
)

var (
	default_avatar = "/static/images/default_avatar.png"
)

func (svc *Service) RegisterUser(c context.Context, req request.ReqisterRequest) (uint32, *errs.ApiError) {
	// 获取验证码
	sms := smsmock.NewSMS(3 * time.Minute)

	smscode, err := sms.GetSMSCode(req.PhoneNumber)
	if err != nil {
		return 0, errs.ErrBadRequest.AsMessage(err.Error()).AsException(err)
	}
	if smscode.Code != req.SMSCode {
		return 0, errs.ErrBadRequest.AsMessage(errs.MsgSMSMismatch).AsException(err)
	}

	hashedPswd, err := pswd.GenHashPassword(req.Password)
	if err != nil {
		return 0, errs.ErrServerError.AsException(err)
	}

	arg := db.CreateUserParams{
		PhoneNumber:    req.PhoneNumber,
		HashedPassword: hashedPswd,
		Username:       req.UserName,
		Avatar:         default_avatar,
	}

	newID, err := service.HandleInsert(c, arg, svc.store.CreateUser)
	if err != nil {
		return 0, errs.HandleSQLError(err)
	}

	return newID, nil
}

// LoginUser 用户登录，成功返回nil, 错误返回error（错误种类：数据库错误，Token错误, ErrMismatchedPaswd，ErrNotImplemented）
func (svc *Service) LoginUser(c context.Context, req request.LoginRequest) (*reps.LoginResponse, *errs.ApiError) {
	user, err := svc.store.GetUserByPhoneNumber(c, req.PhoneNumber)
	if err != nil {
		return nil, errs.HandleSQLError(err)
	}

	if *req.LoginType == request.LoginType_SMS {
		// 验证码登录
		sms := smsmock.NewSMS(3 * time.Minute)
		smscode, err := sms.GetSMSCode(req.PhoneNumber)
		if err != nil {
			return nil, errs.ErrBadRequest.AsMessage(err.Error()).AsException(err)
		}
		if smscode.Code != req.SMSCode {
			return nil, errs.ErrBadRequest.AsMessage(errs.MsgSMSMismatch).AsException(err)
		}
	} else {
		// 密码登录
		err = pswd.CheckPassword(req.Password, user.HashedPassword)
		if err != nil {
			return nil, errs.ErrBadRequest.AsMessage(errs.MsgMismatchedPaswd).AsException(err)
		}
	}

	// 验证成功，做登录的其他业务
	aToken, _, err := configs.TokenMaker.CreateToken(int64(user.ID), configs.Cfg.Token.AccessTokenDuration)
	if err != nil {
		return nil, errs.ErrServerError.AsMessage(errs.MsgCreateTokenFailed).AsException(err)
	}

	rToken, payload, err := configs.TokenMaker.CreateToken(int64(user.ID), configs.Cfg.Token.RefreshTokenDuration)
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
	token, _, err := configs.TokenMaker.CreateToken(int64(session.UserID), configs.Cfg.Token.AccessTokenDuration)
	if err != nil {
		return "", errs.ErrServerError.AsException(err)
	}

	return token, nil
}
