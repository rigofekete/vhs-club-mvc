INSERT INTO users (username, email, role, hashed_password) VALUES
  (
    'Admin', 'admin@vhs-club.hu', 'admin',
    '$argon2id$v=19$m=65536,t=1,p=24$DfvatNddLOcqt5Z0zcKSxg$lnxI8/SEOLUb
81EiCDSN95oaf7MHLvPVv2qgaBevzow'
  ),
  (
    'ArthurCClarke', 'thesentinel@space.odissey', DEFAULT,
    '$argon2id$v=19$m=65536,t=1,p=24$SeQs/E+zWqjrpkfLKWCWNQ$fELoZMYpNKXu
ociqF/RL38OxTh5Zxc97DU0CdY0j3hc'
  ),
  (
    'MilesDavis', 'grumpy.genius@cool.com', DEFAULT,
    '$argon2id$v=19$m=65536,t=1,p=24$SeQs/E+zWqjrpkfLKWCWNQ$fELoZMYpNKXu
ociqF/RL38OxTh5Zxc97DU0CdY0j3hc'
  )
ON CONFLICT DO NOTHING;

INSERT INTO tapes (title, director, genre, quantity, price) VALUES
  ('Amarcord', 'Federico Fellini', 'Drama', 1, 5999.99),
  ('Taxi Driver', 'Martin Scorsese', 'Thriller', 2, 5999.99),
  ('Back to the Future', 'Robert Zemeckis', 'Adventure', 5, 2999.99),
  ('Alien', 'Ridley Scott', 'Horror', 10, 5999.99),
  ('A torinói ló', 'Béla Tarr', 'Drama', 3, 5999.99),
  ('Batman', 'Tim Burton', 'Action', 4, 2999.99),
  ('Fitzcarraldo', 'Werner Herzog', 'Drama', 11, 5999.99)
ON CONFLICT DO NOTHING;
