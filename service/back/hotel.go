package back

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lightsaid/hotel-bookings/api/request"
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
	"github.com/lightsaid/hotel-bookings/service"
)

// CreateHotel 创建酒店
func (svc *Service) CreateHotel(c context.Context, req request.CreateHotelRequest) (uint32, error) {
	arg := db.InsertHotelParams{
		Title:   req.Title,
		Code:    req.Code,
		Address: req.Address,
	}

	return service.HandleInsert(c, arg, svc.store.InsertHotel)
}

// UpdateHotel 更新酒店
func (svc *Service) UpdateHotel(c context.Context, req request.UpdateHotelRequest) error {
	arg := db.UpdateHotelParams{
		Title:   sql.NullString{String: req.Title},
		Code:    sql.NullString{String: req.Code},
		Address: sql.NullString{String: req.Address},
	}

	return svc.store.UpdateHotel(c, arg)
}

// GetListHotels 查询获取酒店列表和总条数
func (svc *Service) GetListHotels(c context.Context, req request.ListRequest) ([]*db.Hotel, int64, error) {
	if req.Keyword == "" {
		arg := db.GetHotelsParams{
			Limit:  req.Limit(),
			Offset: req.Offset(),
		}

		data, err := svc.store.GetHotels(c, arg)
		if err != nil {
			return nil, 0, err
		}

		totalRecords, _ := svc.store.GetHotelsTotalRecords(c, db.GetHotelsTotalRecordsParams(arg))

		return data, totalRecords, nil
	}

	arg := db.GetHotelsByTitleParams{
		Title:  fmt.Sprintf("%s%s%s", "%", req.Keyword, "%"),
		Limit:  req.Limit(),
		Offset: req.Offset(),
	}

	data, err := svc.store.GetHotelsByTitle(c, arg)
	if err != nil {
		return nil, 0, err
	}

	totalRecords, _ := svc.store.GetHotelsByTitleRecords(c, db.GetHotelsByTitleRecordsParams(arg))

	return data, totalRecords, nil
}
