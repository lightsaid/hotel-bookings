package back

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lightsaid/hotel-bookings/api/request"
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
	"github.com/lightsaid/hotel-bookings/pkg/errs"
	"github.com/lightsaid/hotel-bookings/service"
)

// CreateHotel 创建酒店
func (svc *Service) CreateHotel(c context.Context, req request.HotelRequest) (uint32, *errs.ApiError) {
	arg := db.InsertHotelParams{
		Title:   req.Title,
		Code:    req.Code,
		Address: req.Address,
	}

	newID, err := service.HandleInsert(c, arg, svc.store.InsertHotel)
	if err != nil {
		return 0, errs.HandleSQLError(err)
	}

	return newID, nil

}

// UpdateHotel 更新酒店
func (svc *Service) UpdateHotel(c context.Context, req request.HotelRequest) *errs.ApiError {
	arg := db.UpdateHotelParams{
		Title:   sql.NullString{String: req.Title, Valid: true},
		Code:    sql.NullString{String: req.Code, Valid: true},
		Address: sql.NullString{String: req.Address, Valid: true},
		ID:      *req.ID,
	}

	err := svc.store.UpdateHotel(c, arg)
	return errs.HandleSQLError(err)
}

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

// GetHotelByID 获取一个酒店
func (srv *Service) GetHotelByID(c context.Context, id uint32) (*db.Hotel, *errs.ApiError) {
	hotel, err := srv.store.GetHotelByID(c, id)
	if err != nil {
		return nil, errs.HandleSQLError(err)
	}
	return hotel, nil
}

// DeleteHotelByID 删除一个酒店
func (srv *Service) DeleteHotelByID(c context.Context, id uint32) *errs.ApiError {
	err := srv.store.DeleteHotelByID(c, id)
	return errs.HandleSQLError(err)
}

// GetListRoomTypes 获取所有客房类型
func (svc *Service) GetListRoomTypes(c context.Context) ([]*db.RoomType, *errs.ApiError) {
	list, err := svc.store.ListRoomTypes(c)
	return list, errs.HandleSQLError(err)
}

// GetListBookingStatus 获取所有预订状态
func (svc *Service) GetListBookingStatus(c context.Context) ([]*db.BookingStatus, *errs.ApiError) {
	list, err := svc.store.ListBookingStatus(c)
	return list, errs.HandleSQLError(err)
}
