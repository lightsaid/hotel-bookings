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
