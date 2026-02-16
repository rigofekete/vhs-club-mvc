-- +goose Up
CREATE TABLE tapes(
  id          INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  public_id   UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),
  created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
  title       TEXT NOT NULL UNIQUE,
  director    TEXT NOT NULL,
  genre       TEXT NOT NULL,
  quantity    INT NOT NULL,
  price       FLOAT NOT NULL
);

-- +goose Down
DROP TABLE tapes;
