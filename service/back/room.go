package back

import (
	"context"
	"database/sql"

	"github.com/lightsaid/hotel-bookings/api/request"
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
	"github.com/lightsaid/hotel-bookings/service"
)

func (svc *Service) CreateRoom(c context.Context, req request.RoomRequest) (uint32, error) {
	arg := db.InsertRoomParams{
		HotelID:         req.HotelID,
		RoomNumber:      req.RoomNumber,
		RoomImage:       req.RoomImage,
		RoomPrice:       req.RoomPrice,
		BookingStatusID: req.BookingStatusID,
		RoomTypeID:      req.RoomTypeID,
		RoomCapacity:    req.RoomCapacity,
		RoomDescription: req.RoomDescription,
	}

	return service.HandleInsert(c, arg, svc.store.InsertRoom)
}

func (svc *Service) UpdateRoom(c context.Context, req request.RoomRequest) error {
	arg := db.UpdateRoomParams{
		HotelID:         sql.NullInt32{Int32: int32(req.HotelID), Valid: true},
		RoomNumber:      sql.NullString{String: req.RoomNumber, Valid: true},
		RoomImage:       sql.NullString{String: req.RoomImage, Valid: true},
		RoomPrice:       sql.NullInt32{Int32: int32(req.RoomPrice), Valid: true},
		BookingStatusID: sql.NullInt32{Int32: int32(req.BookingStatusID), Valid: true},
		RoomTypeID:      sql.NullInt32{Int32: int32(req.RoomTypeID), Valid: true},
		RoomCapacity:    sql.NullInt32{Int32: int32(req.RoomCapacity), Valid: true},
		RoomDescription: sql.NullString{String: req.RoomDescription, Valid: true},
		ID:              *req.ID,
	}
	return svc.store.UpdateRoom(c, arg)
}

// ListRooms 查询客房列表，如果 req.ID 存在，则查询对应的酒店下的客房
func (svc *Service) ListRooms(c context.Context, req request.ListRequest) (list []*db.Room, total int64, err error) {
	// 没有 hotel id, 查询简单的列表
	if req.ID == nil || *req.ID <= 0 {
		arg := db.GetRoomsParams{
			Limit:  req.Limit(),
			Offset: req.Offset(),
		}

		list, err = svc.store.GetRooms(c, arg)
		if err != nil {
			return
		}
		total, _ = svc.store.GetRoomsTotalRecords(c)
		return
	}

	// 查询对应酒店的客房列表
	arg := db.GetRoomsByHotelIDParams{
		HotelID: *req.ID,
		Limit:   req.Limit(),
		Offset:  req.Offset(),
	}
	list, err = svc.store.GetRoomsByHotelID(c, arg)
	if err != nil {
		return
	}

	total, _ = svc.store.GetRoomsByHotelIDTotalRecords(c, *req.ID)

	return
}

func (svc *Service) GetRoomByID(c context.Context, id uint32) (*db.Room, error) {
	return svc.store.GetRoomByID(c, id)
}

func (svc *Service) DeleteRoom(c context.Context, id uint32) error {
	return svc.store.DeleteRoomByID(c, id)
}

func (svc *Service) UpdateRoomType(c context.Context, req request.UpdateRoomTypeRequest) error {
	arg := db.UpdateRoomTypeParams{
		RoomTypeID: req.RoomTypeID,
		ID:         req.ID,
	}
	return svc.store.UpdateRoomType(c, arg)
}

func (svc *Service) UpdateRoomStatus(c context.Context, req request.UpdateRoomStatusRequest) error {
	arg := db.UpdateRoomBookingStatusParams{
		BookingStatusID: req.BookingStatusID,
		ID:              req.ID,
	}
	return svc.store.UpdateRoomBookingStatus(c, arg)
}
