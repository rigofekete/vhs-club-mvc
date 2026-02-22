-- name: CreateUser :one
INSERT INTO users(username, email, hashed_password)
VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserFromPublicID :one
SELECT * FROM users
WHERE public_id = $1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY created_at ASC;

-- name: DeleteAllUsers :exec
DELETE FROM users;




