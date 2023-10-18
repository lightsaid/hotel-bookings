package back

import (
	"context"

	"github.com/lightsaid/hotel-bookings/api/request"
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
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
