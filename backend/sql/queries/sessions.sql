-- name: CreateSession :exec
INSERT INTO sessions (token, user_id, platform, ip, created_at)
VALUES ( ?, ?, ?, ?, unixepoch('now'));

-- name: GetSession :one
SELECT * FROM sessions
WHERE token = ?;