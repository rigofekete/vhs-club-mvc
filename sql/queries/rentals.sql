-- name: CreateRental :one
INSERT INTO rentals (id, created_at, user_id, tape_id, rented_at, returned_at)
VALUES(
  gen_random_uuid(),
  NOW(),
  $1,
  $2,
  NOW(),
  $3
)
RETURNING *;

-- name: ReturnTape :exec
UPDATE rentals
SET returned_at = $2
WHERE id = $1;

-- name: GetActiveRentalByID :many
SELECT * FROM rentals
-- NULL is not a value so only IS keyword works
WHERE user_id = $1 AND returned_at IS NULL;

-- name: GetActiveRentalbyTape :many
SELECT * FROM rentals
WHERE tape_id = $1 AND returned_at IS NULL;
