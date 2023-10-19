package service

import (
	"context"

	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
	"github.com/lightsaid/hotel-bookings/pkg/errs"
)

type BackService interface {
	CreateHotel(c context.Context, req request.HotelRequest) (uint32, *errs.ApiError)
	UpdateHotel(c context.Context, req request.HotelRequest) *errs.ApiError
	GetListHotels(c context.Context, req request.ListRequest) ([]*db.Hotel, int64, *errs.ApiError)
	GetHotelByID(c context.Context, id uint32) (*db.Hotel, *errs.ApiError)
	DeleteHotelByID(c context.Context, id uint32) *errs.ApiError
	GetListRoomTypes(c context.Context) ([]*db.RoomType, *errs.ApiError)
	GetListBookingStatus(c context.Context) ([]*db.BookingStatus, *errs.ApiError)

	CreateRoom(c context.Context, req request.RoomRequest) (uint32, *errs.ApiError)
	UpdateRoom(c context.Context, req request.RoomRequest) *errs.ApiError
	ListRooms(c context.Context, req request.ListRequest) (list []*db.Room, total int64, err *errs.ApiError)
	GetRoomByID(c context.Context, id uint32) (*db.Room, *errs.ApiError)
	DeleteRoom(c context.Context, id uint32) *errs.ApiError
	UpdateRoomType(c context.Context, req request.UpdateRoomTypeRequest) *errs.ApiError
	UpdateRoomStatus(c context.Context, req request.UpdateRoomStatusRequest) *errs.ApiError

	ListUsers(c context.Context, req request.ListRequest) (list []*db.ListUsersRow, total int64, err *errs.ApiError)
	LoginUser(c context.Context, req request.LoginRequest) (*reps.LoginResponse, *errs.ApiError)
}
