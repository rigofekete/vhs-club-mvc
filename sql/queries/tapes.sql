-- name: CreateTape :one
INSERT INTO tapes (id, created_at, updated_at, title, director, genre, quantity, price)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: GetTapes :many
SELECT * FROM tapes
ORDER BY created_at ASC;

-- name: GetTape :one
SELECT * FROM tapes
WHERE id = $1;

