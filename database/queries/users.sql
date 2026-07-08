-- name: Save :exec
INSERT INTO users (id, name, email, password)
VALUES ($1, $2, $3, $4);

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;
