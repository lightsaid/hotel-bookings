package service

import (
	"context"

	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
	"github.com/lightsaid/hotel-bookings/pkg/errs"
	"github.com/lightsaid/hotel-bookings/pkg/token"
)

type FrontService interface {
	GetListHotels(c context.Context, req request.ListRequest) ([]*db.Hotel, int64, *errs.ApiError)
	QueryRooms(c context.Context, req request.QueryRoomsRequest) ([]*reps.QueryRoomsResponse, int64, *errs.ApiError)

	RegisterUser(c context.Context, req request.ReqisterRequest) (uint32, *errs.ApiError)
	LoginUser(c context.Context, req request.LoginRequest) (*reps.LoginResponse, *errs.ApiError)
	RenewAccessToken(c context.Context, payload *token.Payload, rToken string) (string, *errs.ApiError)
}
