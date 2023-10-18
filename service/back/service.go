package back

import (
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
	"github.com/lightsaid/hotel-bookings/service"
)

type Service struct {
	store db.Store
}

func NewService(store db.Store) service.BackService {
	return &Service{
		store,
	}
}
