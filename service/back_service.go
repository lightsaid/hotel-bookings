package service

import (
	"context"

	"github.com/lightsaid/hotel-bookings/api/request"
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
)

type BackService interface {
	CreateHotel(c context.Context, req request.HotelRequest) (uint32, error)
	UpdateHotel(c context.Context, req request.HotelRequest) error
	GetListHotels(c context.Context, req request.ListRequest) ([]*db.Hotel, int64, error)
	GetHotelByID(c context.Context, id uint32) (*db.Hotel, error)
	DeleteHotelByID(c context.Context, id uint32) error
	GetListRoomTypes(c context.Context) ([]*db.RoomType, error)
	GetListBookingStatus(c context.Context) ([]*db.BookingStatus, error)

	CreateRoom(c context.Context, req request.RoomRequest) (uint32, error)
	UpdateRoom(c context.Context, req request.RoomRequest) error
	ListRooms(c context.Context, req request.ListRequest) (list []*db.Room, total int64, err error)
	GetRoomByID(c context.Context, id uint32) (*db.Room, error)
	DeleteRoom(c context.Context, id uint32) error
	UpdateRoomType(c context.Context, req request.UpdateRoomTypeRequest) error
	UpdateRoomStatus(c context.Context, req request.UpdateRoomStatusRequest) error

	ListUsers(c context.Context, req request.ListRequest) (list []*db.ListUsersRow, total int64, err error)
}
