-- +goose Up
CREATE TABLE rentals(
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  user_id UUID NOT NULL,
  tape_id UUID NOT NULL,
  rented_at TIMESTAMP NOT NULL,
  returned_at TIMESTAMP,
  CONSTRAINT fk_rentals_user
  FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_rentals_tape
  FOREIGN KEY (tape_id) REFERENCES tapes(id)
);

-- +goose Down
DROP TABLE rentals;
