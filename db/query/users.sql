-- name: CreateUser :execresult
INSERT INTO users (
    phone_number,
    hashed_password,
    username,
    avatar
) VALUES (
    ?,?,?,?
);

-- name: UpdateUser :exec
UPDATE 
    users 
SET
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    username = COALESCE(sqlc.narg(username), username),
    avatar = COALESCE(sqlc.narg(avatar), avatar)
WHERE 
    id = ? AND is_deleted = 0;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ? and is_deleted = 0;

-- name: GetUserByPhoneNumber :one
SELECT * FROM users WHERE phone_number = ? and is_deleted = 0;

-- name: ListUsers :many
SELECT 
    id, 
    phone_number, 
    username, 
    avatar, 
    created_at, 
    updated_at
FROM users
WHERE is_deleted = 0
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListUsersTotal :one
SELECT COUNT(*) as total FROM users WHERE is_deleted = 0;