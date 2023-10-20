-- name: CreateSession :execresult
INSERT INTO sessions (
    token_id,
    user_id,
    refresh_token,
    user_agent,
    client_ip,
    expires_at,
    login_type
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
);

-- name: GetSessionByTokenID :one
SELECT * FROM sessions WHERE token_id = ? LIMIT 1;