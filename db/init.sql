-- user table (obviously)
CREATE TABLE users (
  id          BIGSERIAL PRIMARY KEY,
  nickname    TEXT UNIQUE,
  password    TEXT,
  email       TEXT UNIQUE,
  last_seen   TIMESTAMP,
  lat         FLOAT,
  lon         FLOAT,
  token       TEXT,
  picture_url TEXT
);


INSERT INTO
  users (nickname, password, email, last_seen, lat, lon, token, picture_url)
VALUES
  ('Rigfox', 'foobar', 'rigfox@gmail.com', now(), 55.343, 57.432, 'dsad', 'https://i.imgur.com/WixHpwc.jpg'),
  ('Fanzil', 'foobar2', 'kh.fanzil@gmail.com', now(), 12.532, 54.123, '123', 'https://i.imgur.com/DCt8FEC.jpg'),
  ('Mr Foo', 'foobar3', 'foobar@gmail.com', now(), 12.532, 54.123, '47654', 'https://i.imgur.com/dXyI60t.jpg'),
  ('Ms Bar', 'foobar4', 'quxquux@gmail.com', now(), 12.532, 54.123, 'nbvcx', 'https://i.imgur.com/8KjqjzB.jpg'),
  ('Sir Quux', 'foobar5', '2barfoobar@gmail.com', now(), 12.532, 54.123, 'qwry', 'https://i.imgur.com/cGNs3NX.jpg'),
  ('Senior Baz', 'foobar5', '3barfoobar@gmail.com', now(), 12.532, 54.123, 'qwry', 'https://i.imgur.com/XWr5EGP.jpg');

-- user interests
CREATE TABLE user_interests (
  user_id     BIGINT,
  interest_id INT,
  PRIMARY KEY (user_id, interest_id)
);

INSERT INTO
  user_interests
VALUES
  (1, 1), -- sport: football
  (1, 2), -- sport: baseball
  (1, 10), -- music: classical music
  (1, 21), -- IT: fullstack
  (2, 2),
  (2, 6),
  (3, 7),
  (3, 10),
  (4, 3),
  (6, 9);

-- do we really need this? probably not
-- CREATE TABLE interests_categories (
--   id   INT PRIMARY KEY,
--   name TEXT
-- );
--
-- INSERT INTO
--   interests_categories
-- VALUES
--   (1, 'sport'),
--   (2, 'music'),
--   (3, 'outdoor'),
--   (4, 'anime'),
--   (5, 'IT');

CREATE TABLE interests (
  id          SMALLSERIAL PRIMARY KEY,
  category_id INT,
  name        TEXT
);

-- sport interests
INSERT INTO
  interests (category_id, name)
VALUES
  (1, 'football'),
  (1, 'basketball'),
  (1, 'hockey'),
  (1, 'ski'),
  (1, 'snowboard'),
  (1, 'skateboard');

-- music interests
INSERT INTO
  interests (category_id, name)
VALUES
  (2, 'rock'),
  (2, 'pop'),
  (2, 'hip-hop'),
  (2, 'classical music'),
  (2, 'blues'),
  (2, 'jazz'),
  (2, 'electronic');

-- outdoor interests
INSERT INTO
  interests (category_id, name)
VALUES
  (3, 'hiking'),
  (3, 'street walking'),
  (3, 'graffiti');

-- etc

-- anime interests
INSERT INTO
  interests (category_id, name)
VALUES
  (4, 'action'),
  (4, 'drama'),
  (4, 'fantasy'),
  (4, 'sci-fi');

-- IT interests
INSERT INTO
  interests (category_id, name)
VALUES
  (5, 'frontend'),
  (5, 'backend'),
  (5, 'fullstack'),
  (5, 'devops'),
  (5, 'python'),
  (5, 'js'),
  (5, 'java'),
  (5, 'php'),
  (5, 'golang');
-- and more

CREATE TABLE events (
  id          BIGSERIAL PRIMARY KEY,
  name        TEXT,
  description TEXT,
  date        TIMESTAMP,
  owner_id    BIGINT,
  lat         FLOAT,
  lon         FLOAT,
  is_private  BOOLEAN DEFAULT FALSE
);

INSERT INTO
  events (name, description, date, owner_id, lat, lon, is_private)
VALUES
  ('test open event 1', 'description', '2018-03-31 19:10:25-07', 1, 53.927668, 27.685518, FALSE), -- id 1
  ('test open event 2', 'description', '2018-03-12 19:10:25-07', 2, 53.927668, 27.685318, FALSE), -- id 2
  ('test open event 3', 'description', '2010-03-02 19:10:25-07', 2, 53.928568, 27.685618, FALSE), -- id 3
  ('test open event 4', 'description', '2018-01-04 19:10:25-07', 3, 53.928268, 27.685500, FALSE), -- etc
  ('suggested event 1', 'description', '2018-01-04 19:10:25-07', 3, 53.928068, 27.680000, FALSE),
  ('suggested event 2', 'description', '2018-01-04 19:10:25-07', 3, 53.926668, 27.681518, FALSE),
  ('private event 7', 'description', '2018-01-04 19:10:25-07', 4, 53.924668, 27.675518, TRUE),
  ('private event 8', 'description', '2018-01-04 19:10:25-07', 3, 53.929668, 27.687918, TRUE),

  ('test open event 9', 'description', '2018-01-04 19:10:25-07', 2, 53.929168, 27.687918, FALSE),
  ('test open event 10', 'description', '2018-01-04 19:10:25-07', 4, 53.929268, 27.687928, FALSE),
  ('test open event 11', 'description', '2018-01-04 19:10:25-07', 5, 53.929368, 27.687938, FALSE),
  ('test open event 12', 'description', '2018-01-04 19:10:25-07', 5, 53.929468, 27.687948, FALSE),
  ('test open event 13', 'description', '2018-01-04 19:10:25-07', 5, 53.929568, 27.687958, FALSE),
  ('test open event 14', 'description', '2018-01-04 19:10:25-07', 4, 53.929668, 27.687968, FALSE);

-- event tags
CREATE TABLE event_interests (
  event_id    BIGINT,
  interest_id BIGINT,
  PRIMARY KEY (event_id, interest_id)
);

INSERT INTO
  event_interests
VALUES
  (1, 1), -- sport: football
  (1, 2), -- sport: baseball
  (5, 1), -- sport: football
  (5, 2), -- sport: baseball
  (5, 10), -- music: classical music
  (6, 12), -- music: jazz
  (6, 17), -- films: action
  (6, 20), -- films: sci-fi
  (6, 9), -- music: hip-hop
  (6, 4), -- sport: ski
  (6, 21), -- sport: ski
  (13, 1), -- sport: ski
  (13, 2), -- sport: ski
  (13, 4), -- sport: ski
  (13, 5); -- sport: ski

--- event attendees
CREATE TABLE event_attendees (
  event_id BIGINT,
  user_id  BIGINT,
  PRIMARY KEY (event_id, user_id)
);

INSERT INTO
  event_attendees
VALUES
  (1, 1),
  (2, 1),
  (5, 1),
  (7, 1),

  (1, 3),
  (2, 2),
  (2, 3),
  (3, 3),
  (4, 3);


