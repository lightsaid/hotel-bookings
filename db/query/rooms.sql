-- name: InsertRoom :execresult
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
);

-- name: UpdateRoom :exec
UPDATE rooms 
SET
    hotel_id = COALESCE(sqlc.narg(hotel_id), hotel_id),
    room_number = COALESCE(sqlc.narg(room_number), room_number),
    room_image = COALESCE(sqlc.narg(room_image), room_image),
    room_price = COALESCE(sqlc.narg(room_price), room_price),
    booking_status_id = COALESCE(sqlc.narg(booking_status_id), booking_status_id),
    room_type_id = COALESCE(sqlc.narg(room_type_id), room_type_id),
    room_capacity = COALESCE(sqlc.narg(room_capacity), room_capacity),
    room_description = COALESCE(sqlc.narg(room_description), room_description),
    updated_at = now()
WHERE id = sqlc.arg(id) AND is_deleted = 0;


-- name: GetRooms :many
SELECT * FROM rooms WHERE is_deleted = 0 ORDER BY updated_at LIMIT ? OFFSET ?;

-- name: GetRoomsTotalRecords :one
SELECT COUNT(*) as total_records FROM rooms WHERE is_deleted = 0;

-- name: GetRoomsByHotelID :many
SELECT * FROM rooms WHERE hotel_id = ? AND is_deleted = 0 ORDER BY updated_at LIMIT ? OFFSET ?;

-- name: GetRoomsByHotelIDTotalRecords :one
SELECT COUNT(*) as total_records FROM rooms WHERE hotel_id = ? AND is_deleted = 0;

-- name: GetRoomByID :one
SELECT * FROM rooms WHERE id = ? AND is_deleted = 0;

-- name: DeleteRoomByID :exec
UPDATE rooms SET is_deleted = 1, updated_at = NOW() WHERE id = ? AND is_deleted = 0;

-- name: UpdateRoomBookingStatus :exec
UPDATE rooms SET booking_status_id = ?, updated_at = NOW() WHERE id = ? AND is_deleted = 0;

-- name: UpdateRoomType :exec
UPDATE rooms SET room_type_id = ?, updated_at = NOW() WHERE id = ? AND is_deleted = 0;

-- name: QueryRooms :many
SELECT 
    r.id,
    r.hotel_id,
    r.room_number,
    r.room_image,
    r.room_price,
    r.booking_status_id,
    r.room_type_id,
    r.room_capacity,
    r.room_description,
    h.title,
    h.code,
    rt.room_label,
    rt.room_type,
    bs.status_label,
    bs.status_type
FROM 
    rooms r  
INNER JOIN 
    hotels h
ON 
    h.id = r.hotel_id
LEFT JOIN 
    room_types rt
ON 
    rt.id = r.room_type_id
LEFT JOIN
    booking_status bs
ON
    bs.id = r.booking_status_id
WHERE 
    r.is_deleted = 0
    AND (CASE WHEN sqlc.arg(hotel_id) > 0 THEN r.hotel_id = sqlc.arg(hotel_id) ELSE TRUE END)
    AND (CASE WHEN sqlc.arg(room_type_id) > 0 THEN r.hotel_id = sqlc.arg(room_type_id) ELSE TRUE END)
    AND (CASE WHEN sqlc.arg(booking_status_id) > 0 THEN r.booking_status_id = sqlc.arg(booking_status_id) ELSE TRUE END)
ORDER BY r.id 
LIMIT ? OFFSET ?;

-- name: QueryRoomsTotal :one
SELECT COUNT(*) as total
FROM 
    rooms r  
INNER JOIN 
    hotels h
ON 
    h.id = r.hotel_id
LEFT JOIN 
    room_types rt
ON 
    rt.id = r.room_type_id
LEFT JOIN
    booking_status bs
ON
    bs.id = r.booking_status_id
WHERE 
    r.is_deleted = 0
    AND (CASE WHEN sqlc.arg(hotel_id) > 0 THEN r.hotel_id = sqlc.arg(hotel_id) ELSE TRUE END)
    AND (CASE WHEN sqlc.arg(room_type_id) > 0 THEN r.hotel_id = sqlc.arg(room_type_id) ELSE TRUE END)
    AND (CASE WHEN sqlc.arg(booking_status_id) > 0 THEN r.booking_status_id = sqlc.arg(booking_status_id) ELSE TRUE END)