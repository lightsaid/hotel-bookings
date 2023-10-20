-- name: CreateBookings :execresult
INSERT INTO bookings (
    id,
    user_id,
    booking_from,
    booking_to,
    room_id,
    members,
    total_amount
) VALUES (
    ?,?,?,?,?,?,?
);

-- name: ListBookings :many
SELECT * FROM bookings WHERE is_deleted = 0 LIMIT ? OFFSET ?;

-- name: ListBookingsByUserID :many
SELECT * FROM bookings WHERE is_deleted = 0 AND user_id = ? LIMIT ? OFFSET ?;

-- name: DeleteBookings :exec
UPDATE bookings SET is_deleted = 1, updated_at = NOW() WHERE id = ? AND is_deleted = 0;

