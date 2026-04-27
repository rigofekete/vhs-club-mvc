-- name: CreateTape :one
INSERT INTO tapes (title, director, genre, quantity)
VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING *;

-- name: GetTapes :many
SELECT * FROM tapes
ORDER BY created_at ASC;

-- name: GetTapeByID :one
SELECT * FROM tapes
WHERE id = $1;

-- name: GetTapeFromPublicID :one
SELECT * FROM tapes
WHERE public_id = $1;

-- name: UpdateTape :one
UPDATE tapes
SET
  updated_at =  NOW(),
  title =       COALESCE(sqlc.narg('title'), title),
  director =    COALESCE(sqlc.narg('director'), director),
  genre =       COALESCE(sqlc.narg('genre'), genre),
  quantity =    COALESCE(sqlc.narg('quantity'), quantity)
WHERE id = $1
RETURNING *;

-- name: DeleteTape :exec
DELETE FROM tapes
WHERE id = $1;

-- name: DeleteAllTapes :exec
DELETE FROM tapes;
