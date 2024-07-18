-- name: GetUser :one
SELECT * FROM users
WHERE id = ?;

-- name: CheckUsername :one
SELECT id FROM users
WHERE username = ?;

-- name: CheckEmail :one
SELECT id FROM users
WHERE email = ?;

-- name: CreateUser :exec
INSERT INTO users (id, username, email, password, created_at, updated_at)
VALUES ( ?, ?, ?, ?, unixepoch('now'), unixepoch('now'));