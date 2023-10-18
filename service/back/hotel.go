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
func (svc *Service) CreateHotel(c context.Context, req request.HotelRequest) (uint32, error) {
	arg := db.InsertHotelParams{
		Title:   req.Title,
		Code:    req.Code,
		Address: req.Address,
	}

	return service.HandleInsert(c, arg, svc.store.InsertHotel)
}

// UpdateHotel 更新酒店
func (svc *Service) UpdateHotel(c context.Context, req request.HotelRequest) error {
	arg := db.UpdateHotelParams{
		Title:   sql.NullString{String: req.Title, Valid: true},
		Code:    sql.NullString{String: req.Code, Valid: true},
		Address: sql.NullString{String: req.Address, Valid: true},
		ID:      *req.ID,
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
		return nil, 0, err
	}

	totalRecords, _ := svc.store.GetHotelsByTitleRecords(c, arg.Title)

	return data, totalRecords, nil
}

// GetHotelByID 获取一个酒店
func (srv *Service) GetHotelByID(c context.Context, id uint32) (*db.Hotel, error) {
	return srv.store.GetHotelByID(c, id)
}

// DeleteHotelByID 删除一个酒店
func (srv *Service) DeleteHotelByID(c context.Context, id uint32) error {
	return srv.store.DeleteHotelByID(c, id)
}

// GetListRoomTypes 获取所有客房类型
func (svc *Service) GetListRoomTypes(c context.Context) ([]*db.RoomType, error) {
	return svc.store.ListRoomTypes(c)
}

// GetListBookingStatus 获取所有预订状态
func (svc *Service) GetListBookingStatus(c context.Context) ([]*db.BookingStatus, error) {
	return svc.store.ListBookingStatus(c)
}
