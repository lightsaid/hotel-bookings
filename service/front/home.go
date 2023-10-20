package front

import (
	"context"
	"fmt"

	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
	"github.com/lightsaid/hotel-bookings/pkg/errs"
)

// GetListHotels 查询获取酒店列表和总条数
func (svc *Service) GetListHotels(c context.Context, req request.ListRequest) ([]*db.Hotel, int64, *errs.ApiError) {
	if req.Keyword == "" {
		arg := db.GetHotelsParams{
			Limit:  req.Limit(),
			Offset: req.Offset(),
		}

		data, err := svc.store.GetHotels(c, arg)
		if err != nil {
			return nil, 0, errs.HandleSQLError(err)
		}

		totalRecords, _ := svc.store.GetHotelsTotalRecords(c)

		return data, totalRecords, nil
	}

	arg := db.GetHotelsByTitleParams{
		Title:  fmt.Sprintf("%s%s%s", "%", req.Keyword, "%"),
		Limit:  req.Limit(),
		Offset: req.Offset(),
	}

	data, err := svc.store.GetHotelsByTitle(c, arg)
	if err != nil {
		return nil, 0, errs.HandleSQLError(err)
	}

	totalRecords, _ := svc.store.GetHotelsByTitleRecords(c, arg.Title)

	return data, totalRecords, nil
}

// QueryRooms 根据条件查询客房列表
func (svc *Service) QueryRooms(c context.Context, req request.QueryRoomsRequest) ([]*reps.QueryRoomsResponse, int64, *errs.ApiError) {
	arg := db.QueryRoomsParams{
		Limit:           req.Limit(),
		Offset:          req.Offset(),
		HotelID:         req.HotelID,
		RoomTypeID:      req.RoomTypeID,
		BookingStatusID: req.BookingStatusID,
	}

	list, err := svc.store.QueryRooms(c, arg)
	if err != nil {
		return nil, 0, errs.HandleSQLError(err)
	}

	totalArg := db.QueryRoomsTotalParams{
		HotelID:         req.HotelID,
		RoomTypeID:      req.RoomTypeID,
		BookingStatusID: req.BookingStatusID,
	}

	total, _ := svc.store.QueryRoomsTotal(c, totalArg)

	var datas = make([]*reps.QueryRoomsResponse, len(list))
	for i := range list {
		datas[i] = reps.ToQueryRoomsResponse(list[i])
	}

	return datas, total, nil
}
