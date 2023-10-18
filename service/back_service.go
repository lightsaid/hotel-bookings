package service

import (
	"context"

	"github.com/lightsaid/hotel-bookings/api/request"
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
)

type BackService interface {
	CreateHotel(c context.Context, req request.CreateHotelRequest) (uint32, error)
	UpdateHotel(c context.Context, req request.UpdateHotelRequest) error
	GetListHotels(c context.Context, req request.ListRequest) ([]*db.Hotel, int64, error)
}
