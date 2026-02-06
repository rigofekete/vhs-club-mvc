-- +goose Up
CREATE TABLE tapes(
  id UUID     PRIMARY KEY,
  created_at  TIMESTAMP NOT NULL,
  updated_at  TIMESTAMP NOT NULL,
  title       TEXT NOT NULL,
  director    TEXT NOT NULL,
  genre       TEXT NOT NULL,
  quantity    INT NOT NULL,
  price       INT NOT NULL
  -- user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
