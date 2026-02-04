-- name: GetAllUsers :many
SELECT * from users;

-- name: RegisterUser :one
INSERT INTO users (name, email, password, email_verified_at) values($1, $2, $3, NULL) RETURNING *;

-- name: GetUserByToken
