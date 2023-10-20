package request

type RoomRequest struct {
	RoomNumber      string  `json:"room_number" binding:"required"`
	RoomImage       string  `json:"room_image" binding:"required"`
	RoomPrice       uint32  `json:"room_price" binding:"required,min=0"`
	BookingStatusID uint32  `json:"booking_status_id" binding:"required,min=1"`
	RoomTypeID      uint32  `json:"room_type_id" binding:"required,min=1"`
	RoomCapacity    uint32  `json:"room_capacity"  binding:"required,min=1"`
	RoomDescription string  `json:"room_description"`
	HotelID         uint32  `json:"hotel_id"`
	ID              *uint32 `json:"id,omitempty" binding:"omitempty,min=1"`
}

type UpdateRoomTypeRequest struct {
	ID         uint32 `json:"id" binding:"required,min=1"`
	RoomTypeID uint32 `json:"room_type_id" binding:"required,min=1"`
}

type UpdateRoomStatusRequest struct {
	ID              uint32 `json:"id" binding:"required,min=1"`
	BookingStatusID uint32 `json:"booking_status_id" binding:"required,min=1"`
}

type QueryRoomsRequest struct {
	PageNum         int32  `json:"page_num" binding:"required,min=1"`
	PageSize        int32  `json:"page_size" binding:"required,min=5"`
	HotelID         uint32 `json:"hotel_id"`
	BookingStatusID uint32 `json:"booking_status_id"`
	RoomTypeID      uint32 `json:"room_type_id"`
}

func (req QueryRoomsRequest) Limit() int32 {
	if req.PageSize > 100 {
		req.PageSize = 100
	}
	if req.PageSize < 5 {
		req.PageSize = 5
	}
	return req.PageSize
}

func (req QueryRoomsRequest) Offset() int32 {
	return (req.PageNum - 1) * req.Limit()
}
