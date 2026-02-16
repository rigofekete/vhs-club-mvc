-- +goose Up
CREATE TABLE rentals(
  id            INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  public_id     UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),
  created_at    TIMESTAMP NOT NULL DEFAULT NOW(),
  user_id       INT NOT NULL,
  tape_id       INT NOT NULL,
  rented_at     TIMESTAMP NOT NULL DEFAULT NOW(),
  returned_at   TIMESTAMP,
  CONSTRAINT fk_rentals_user
  FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_rentals_tape
  FOREIGN KEY (tape_id) REFERENCES tapes(id)
);

-- +goose Down
DROP TABLE rentals;
