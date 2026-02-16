-- name: CreateTape :one
INSERT INTO tapes (title, director, genre, quantity, price)
VALUES (
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

-- name: GetTapeByID :one
SELECT * FROM tapes
WHERE id = $1;

-- name: GetTapeByPublicID :one
SELECT * FROM tapes
WHERE public_id = $1;

-- name: UpdateTape :one
UPDATE tapes
SET
  updated_at =  NOW(),
  title =       COALESCE(sqlc.narg('title'), title),
  director =    COALESCE(sqlc.narg('director'), director),
  genre =       COALESCE(sqlc.narg('genre'), genre),
  quantity =    COALESCE(sqlc.narg('quantity'), quantity),
  price =       COALESCE(sqlc.narg('price'), price)
WHERE id = $1
RETURNING *;

-- name: DeleteTape :exec
DELETE FROM tapes
WHERE id = $1;

-- name: DeleteAllTapes :exec
DELETE FROM tapes;
