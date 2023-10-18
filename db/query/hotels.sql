-- name: InsertHotel :execresult
INSERT INTO hotels (title, code, address) VALUES (?, ?, ?);

-- name: UpdateHotel :exec
UPDATE hotels
SET
    title = COALESCE(sqlc.narg(title), title),
    code = COALESCE(sqlc.narg(code), code),
    address = COALESCE(sqlc.narg(address), address),
    updated_at = NOW()
WHERE id = sqlc.arg(id) AND is_deleted = 0;

-- name: GetHotelByID :one
SELECT * FROM hotels WHERE id = ? AND is_deleted = 0;

-- name: GetHotels :many
SELECT * FROM hotels WHERE is_deleted = 0 ORDER BY id LIMIT ? OFFSET ?;

-- name: GetHotelsTotalRecords :one
SELECT COUNT(*) as total_records FROM hotels WHERE is_deleted = 0;

-- name: GetHotelsByTitle :many
SELECT * FROM hotels WHERE title LIKE ? AND is_deleted = 0 ORDER BY id LIMIT ? OFFSET ?;

-- name: GetHotelsByTitleRecords :one
SELECT COUNT(*) as total_records FROM hotels WHERE title LIKE ? AND is_deleted = 0;

-- name: DeleteHotelByID :exec
UPDATE hotels SET is_deleted = 1, updated_at = NOW() WHERE id = ? and is_deleted = 0;

