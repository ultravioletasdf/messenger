-- name: GetUser :one
SELECT * FROM users
WHERE id = ?;

-- name: CreateUser :exec
INSERT INTO users (id, username, email, password, created_at, updated_at)
VALUES ( ?, ?, ?, ?, unixepoch('now'), unixepoch('now'));