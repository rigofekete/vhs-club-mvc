-- +goose Up
CREATE TABLE users(
  id          INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  public_id   UUID UNIQUE DEFAULT gen_random_uuid(),
  created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
  username    TEXT  NOT NULL UNIQUE,
  email       TEXT  NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE users;
