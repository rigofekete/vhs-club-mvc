CREATE DATABASE vhs_club;

\c vhs_club

CREATE TABLE users (
  id               INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  public_id        UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),
  created_at       TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at       TIMESTAMP NOT NULL DEFAULT NOW(),
  username         TEXT  NOT NULL UNIQUE,
  email            TEXT  NOT NULL UNIQUE,
  role             TEXT NOT NULL DEFAULT 'user',
  hashed_password  TEXT NOT NULL
);

INSERT INTO users (username, email, role, hashed_password) VALUES
  (
    'Admin', 'admin@vhs-club.hu', 'admin',
    '$argon2id$v=19$m=65536,t=1,p=24$DfvatNddLOcqt5Z0zcKSxg$lnxI8/SEOLUb81EiCDSN95oaf7MHLvPVv2qgaBevzow'
  ),
  (
    'ArthurCClarke', 'thesentinel@space.odissey', DEFAULT,
    '$argon2id$v=19$m=65536,t=1,p=24$SeQs/E+zWqjrpkfLKWCWNQ$fELoZMYpNKXuociqF/RL38OxTh5Zxc97DU0CdY0j3hc'
  ),
  (
    'MilesDavis', 'grumpy.genius@cool.com', DEFAULT,
    '$argon2id$v=19$m=65536,t=1,p=24$SeQs/E+zWqjrpkfLKWCWNQ$fELoZMYpNKXuociqF/RL38OxTh5Zxc97DU0CdY0j3hc'
  );


CREATE TABLE tapes (
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

INSERT INTO tapes (title, director, genre, quantity, price) VALUES
  ('Amarcord', 'Federico Fellini', 'Drama', 1, 5999.99),
  ('Taxi Driver', 'Martin Scorsese', 'Thriller', 2, 5999.99),
  ('Back to the Future', 'Robert Zemeckis', 'Adventure', 5, 2999.99),
  ('Alien', 'Ridley Scott', 'Horror', 10, 5999.99),
  ('A torinói ló', 'Béla Tarr', 'Drama', 3, 5999.99),
  ('Batman', 'Tim Burton', 'Action', 4, 2999.99),
  ('Fitzcarraldo', 'Werner Herzog', 'Drama', 11, 5999.99);

CREATE TABLE rentals (
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



