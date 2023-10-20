package reps

import db "github.com/lightsaid/hotel-bookings/db/sqlc"

type QueryRoomsResponse struct {
	ID              uint32 `json:"id"`
	HotelID         uint32 `json:"hotel_id"`
	RoomNumber      string `json:"room_number"`
	RoomImage       string `json:"room_image"`
	RoomPrice       uint32 `json:"room_price"`
	BookingStatusID uint32 `json:"booking_status_id"`
	RoomTypeID      uint32 `json:"room_type_id"`
	RoomCapacity    uint32 `json:"room_capacity"`
	RoomDescription string `json:"room_description"`
	Title           string `json:"title"` // 酒店 title
	Code            string `json:"code"`  // 酒店 code
	RoomLabel       string `json:"room_label"`
	RoomType        string `json:"room_type"`
	StatusLabel     string `json:"status_label"`
	StatusType      string `json:"status_type"`
}

func ToQueryRoomsResponse(row *db.QueryRoomsRow) *QueryRoomsResponse {
	return &QueryRoomsResponse{
		ID:              row.ID,
		HotelID:         row.HotelID,
		RoomNumber:      row.RoomNumber,
		RoomImage:       row.RoomImage,
		RoomPrice:       row.RoomPrice,
		BookingStatusID: row.BookingStatusID,
		RoomTypeID:      row.RoomTypeID,
		RoomCapacity:    row.RoomCapacity,
		RoomDescription: row.RoomDescription,
		Title:           row.Title,
		Code:            row.Code,
		RoomLabel:       row.RoomLabel.String,
		RoomType:        row.RoomType.String,
		StatusLabel:     row.StatusLabel.String,
		StatusType:      row.StatusType.String,
	}
}
