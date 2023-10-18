// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: rooms.sql

package db

import (
	"context"
	"database/sql"
)

const DeleteRoomByID = `-- name: DeleteRoomByID :exec
UPDATE rooms SET is_deleted = 1, updated_at = NOW() WHERE id = ? AND is_deleted = 0
`

func (q *Queries) DeleteRoomByID(ctx context.Context, id uint32) error {
	_, err := q.db.ExecContext(ctx, DeleteRoomByID, id)
	return err
}

const GetRoomByID = `-- name: GetRoomByID :one
SELECT id, hotel_id, room_number, room_image, room_price, booking_status_id, room_type_id, room_capacity, room_description, created_at, updated_at, is_deleted FROM rooms WHERE id = ? AND is_deleted = 0
`

func (q *Queries) GetRoomByID(ctx context.Context, id uint32) (*Room, error) {
	row := q.db.QueryRowContext(ctx, GetRoomByID, id)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.HotelID,
		&i.RoomNumber,
		&i.RoomImage,
		&i.RoomPrice,
		&i.BookingStatusID,
		&i.RoomTypeID,
		&i.RoomCapacity,
		&i.RoomDescription,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
	)
	return &i, err
}

const GetRooms = `-- name: GetRooms :many
SELECT id, hotel_id, room_number, room_image, room_price, booking_status_id, room_type_id, room_capacity, room_description, created_at, updated_at, is_deleted FROM rooms WHERE is_deleted = 0 ORDER BY updated_at LIMIT ? OFFSET ?
`

type GetRoomsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetRooms(ctx context.Context, arg GetRoomsParams) ([]*Room, error) {
	rows, err := q.db.QueryContext(ctx, GetRooms, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Room{}
	for rows.Next() {
		var i Room
		if err := rows.Scan(
			&i.ID,
			&i.HotelID,
			&i.RoomNumber,
			&i.RoomImage,
			&i.RoomPrice,
			&i.BookingStatusID,
			&i.RoomTypeID,
			&i.RoomCapacity,
			&i.RoomDescription,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsDeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetRoomsByHotelID = `-- name: GetRoomsByHotelID :many
SELECT id, hotel_id, room_number, room_image, room_price, booking_status_id, room_type_id, room_capacity, room_description, created_at, updated_at, is_deleted FROM rooms WHERE hotel_id = ? AND is_deleted = 0 ORDER BY updated_at LIMIT ? OFFSET ?
`

type GetRoomsByHotelIDParams struct {
	HotelID uint32 `json:"hotel_id"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

func (q *Queries) GetRoomsByHotelID(ctx context.Context, arg GetRoomsByHotelIDParams) ([]*Room, error) {
	rows, err := q.db.QueryContext(ctx, GetRoomsByHotelID, arg.HotelID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Room{}
	for rows.Next() {
		var i Room
		if err := rows.Scan(
			&i.ID,
			&i.HotelID,
			&i.RoomNumber,
			&i.RoomImage,
			&i.RoomPrice,
			&i.BookingStatusID,
			&i.RoomTypeID,
			&i.RoomCapacity,
			&i.RoomDescription,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsDeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetRoomsByHotelIDTotalRecords = `-- name: GetRoomsByHotelIDTotalRecords :one
SELECT COUNT(*) as total_records FROM rooms WHERE hotel_id = ? AND is_deleted = 0
`

func (q *Queries) GetRoomsByHotelIDTotalRecords(ctx context.Context, hotelID uint32) (int64, error) {
	row := q.db.QueryRowContext(ctx, GetRoomsByHotelIDTotalRecords, hotelID)
	var total_records int64
	err := row.Scan(&total_records)
	return total_records, err
}

const GetRoomsTotalRecords = `-- name: GetRoomsTotalRecords :one
SELECT COUNT(*) as total_records FROM rooms WHERE is_deleted = 0
`

func (q *Queries) GetRoomsTotalRecords(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, GetRoomsTotalRecords)
	var total_records int64
	err := row.Scan(&total_records)
	return total_records, err
}

const InsertRoom = `-- name: InsertRoom :execresult
INSERT INTO rooms(
    hotel_id,
    room_number,
    room_image,
    room_price,
    booking_status_id,
    room_type_id,
    room_capacity,
    room_description
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
)
`

type InsertRoomParams struct {
	HotelID         uint32 `json:"hotel_id"`
	RoomNumber      string `json:"room_number"`
	RoomImage       string `json:"room_image"`
	RoomPrice       uint32 `json:"room_price"`
	BookingStatusID uint32 `json:"booking_status_id"`
	RoomTypeID      uint32 `json:"room_type_id"`
	RoomCapacity    uint32 `json:"room_capacity"`
	RoomDescription string `json:"room_description"`
}

func (q *Queries) InsertRoom(ctx context.Context, arg InsertRoomParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, InsertRoom,
		arg.HotelID,
		arg.RoomNumber,
		arg.RoomImage,
		arg.RoomPrice,
		arg.BookingStatusID,
		arg.RoomTypeID,
		arg.RoomCapacity,
		arg.RoomDescription,
	)
}

const UpdateRoom = `-- name: UpdateRoom :exec
UPDATE rooms 
SET
    hotel_id = COALESCE(?, hotel_id),
    room_number = COALESCE(?, room_number),
    room_image = COALESCE(?, room_image),
    room_price = COALESCE(?, room_price),
    booking_status_id = COALESCE(?, booking_status_id),
    room_type_id = COALESCE(?, room_type_id),
    room_capacity = COALESCE(?, room_capacity),
    room_description = COALESCE(?, room_description),
    updated_at = now()
WHERE id = ? AND is_deleted = 0
`

type UpdateRoomParams struct {
	HotelID         sql.NullInt32  `json:"hotel_id"`
	RoomNumber      sql.NullString `json:"room_number"`
	RoomImage       sql.NullString `json:"room_image"`
	RoomPrice       sql.NullInt32  `json:"room_price"`
	BookingStatusID sql.NullInt32  `json:"booking_status_id"`
	RoomTypeID      sql.NullInt32  `json:"room_type_id"`
	RoomCapacity    sql.NullInt32  `json:"room_capacity"`
	RoomDescription sql.NullString `json:"room_description"`
	ID              uint32         `json:"id"`
}

func (q *Queries) UpdateRoom(ctx context.Context, arg UpdateRoomParams) error {
	_, err := q.db.ExecContext(ctx, UpdateRoom,
		arg.HotelID,
		arg.RoomNumber,
		arg.RoomImage,
		arg.RoomPrice,
		arg.BookingStatusID,
		arg.RoomTypeID,
		arg.RoomCapacity,
		arg.RoomDescription,
		arg.ID,
	)
	return err
}

const UpdateRoomBookingStatus = `-- name: UpdateRoomBookingStatus :exec
UPDATE rooms SET booking_status_id = ?, updated_at = NOW() WHERE id = ? AND is_deleted = 0
`

type UpdateRoomBookingStatusParams struct {
	BookingStatusID uint32 `json:"booking_status_id"`
	ID              uint32 `json:"id"`
}

func (q *Queries) UpdateRoomBookingStatus(ctx context.Context, arg UpdateRoomBookingStatusParams) error {
	_, err := q.db.ExecContext(ctx, UpdateRoomBookingStatus, arg.BookingStatusID, arg.ID)
	return err
}

const UpdateRoomType = `-- name: UpdateRoomType :exec
UPDATE rooms SET room_type_id = ?, updated_at = NOW() WHERE id = ? AND is_deleted = 0
`

type UpdateRoomTypeParams struct {
	RoomTypeID uint32 `json:"room_type_id"`
	ID         uint32 `json:"id"`
}

func (q *Queries) UpdateRoomType(ctx context.Context, arg UpdateRoomTypeParams) error {
	_, err := q.db.ExecContext(ctx, UpdateRoomType, arg.RoomTypeID, arg.ID)
	return err
}
