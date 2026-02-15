-- name: CreateUser :one
INSERT INTO users(username, email)
VALUES (
  $1,
  $2
)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY created_at ASC;

-- name: DeleteAllUsers :exec
DELETE FROM users;




