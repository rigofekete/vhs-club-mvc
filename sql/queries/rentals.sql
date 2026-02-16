-- name: CreateRental :one
INSERT INTO rentals (user_id, tape_id)
VALUES(
  $1,
  $2
)
RETURNING *;

-- name: ReturnTape :exec
UPDATE rentals
SET returned_at = $2
WHERE id = $1;

-- name: GetActiveRentalByUserID :many
SELECT * FROM rentals
-- NULL is not a value so only IS keyword works
WHERE user_id = $1 AND returned_at IS NULL;

-- name: GetActiveRentalbyTape :many
SELECT * FROM rentals
WHERE tape_id = $1 AND returned_at IS NULL;

-- name: GetActiveRentalCountByTape :one
SELECT COUNT(*) FROM rentals
WHERE tape_id = $1 AND returned_at IS NULL;

-- name: GetAllActiveRentals :many
SELECT * FROM rentals
WHERE returned_at IS NULL
ORDER BY created_at ASC;
