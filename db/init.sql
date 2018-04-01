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
  ('Rigfox', 'foobar', 'rigfox@gmail.com', now(), 55.343, 57.432, 'dsad', 'https://i.imgur.com/cGNs3NX.jpg'),
  ('Fanzil', 'foobar2', 'kh.fanzil@gmail.com', now(), 12.532, 54.123, '123', 'https://i.imgur.com/DCt8FEC.jpg'),
  ('Mr Foo', 'foobar3', 'foobar@gmail.com', now(), 12.532, 54.123, '47654', 'https://i.imgur.com/dXyI60t.jpg'),
  ('Potato Party Bot', 'foobar4', 'quxquux@gmail.com', now(), 12.532, 54.123, 'nbvcx', 'https://i.imgur.com/8KjqjzB.jpg'),
  ('Sir Quux', 'foobar5', '2barfoobar@gmail.com', now(), 12.532, 54.123, 'qwry', 'https://i.imgur.com/WixHpwc.jpg'),
  ('Erlene Villano', 'foobar5', '3barfoobar@gmail.com', now(), 12.532, 54.123, 'qwry', 'https://i.imgur.com/XWr5EGP.jpg'),
  ('Amber Rathbone', 'foobar5', '1@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Vasiliki Daum', 'foobar5', '2@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Brittany Sexton', 'foobar5', '3@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Dannie Hinkle', 'foobar5', '4@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Vera Oboyle', 'foobar5', '5@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Kerry Kies', 'foobar5', '6@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Naida Nyberg', 'foobar5', '7@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Cammy Calton', 'foobar5', '8@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Ginger Gatto', 'foobar5', '8xx@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Debby Dyck', 'foobar5', '9@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Jasmin Jamar', 'foobar5', 'q@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Lizzie Licari', 'foobar5', 'w@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Angila Angstadt', 'foobar5', 'e@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Evalyn Erskine', 'foobar5', 'r@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Cleora Cutlip', 'foobar5', 't@random.com', now(), 12.532, 54.123, 'qwry', 'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Jadwiga Jumper', 'foobar5', 'y@random.com', now(), 12.532, 54.123, 'qwry',
   'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Shanna Silvis', 'foobar5', 'u@random.com', now(), 12.532, 54.123, 'qwry',
   'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Markus Mullaney', 'foobar5', 's@random.com', now(), 12.532, 54.123, 'qwry',
   'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Karin Kulpa', 'foobar5', 'd@random.com', now(), 12.532, 54.123, 'qwry',
   'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg'),
  ('Alfonso Arbogast', 'foobar5', 'd2@random.com', now(), 12.532, 54.123, 'qwry',
   'https://sun9-7.userapi.com/c824501/v824501404/7d3e5/Pd0pYzunjDU.jpg');

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
  events (name, description, date, owner_id, is_private, lat, lon)
VALUES
  ('Смотрим ЧМ 2018!', 'Давайте сегодня соберёмся посмотреть матч...', '2018-03-31 19:10:25-07', 1, FALSE, 53.927220, 27.682033),
  ('Клуб слоу-мо чуваков', 'Не опаздывайте', '2018-03-12 19:10:25-07', 2, FALSE, 53.928483, 27.692590),
  ('Дегустация кексов здесь', 'Кексы здесь', '2018-01-04 19:10:25-07', 3, FALSE, 53.927332, 27.688273),
  ('Хакатоним!', 'Хакатоны для про', '2018-01-04 19:10:25-07', 3, FALSE, 53.928261, 27.685676),
  ('Калик го пыхать', 'Калик топ тема', '2018-01-04 19:10:25-07', 4, TRUE, 53.928161, 27.683420),
  ('Играем в ГТА 5', 'description', '2018-01-04 19:10:25-07', 3, TRUE, 53.925788, 27.686430),
  ('Го бухать, ёпта', 'description', '2018-01-04 19:10:25-07', 2, FALSE, 53.927295, 27.688148),
  ('Говнокодим в баре', 'description', '2018-01-04 19:10:25-07', 5, FALSE, 53.929147, 27.685999),
  ('Сдаём диплом', 'description', '2018-01-04 19:10:25-07', 4, TRUE, 53.928225, 27.688225);

-- event tags
CREATE TABLE event_interests (
  event_id    BIGINT,
  interest_id BIGINT,
  PRIMARY KEY (event_id, interest_id)
);

-- INSERT INTO
--   event_interests
-- VALUES
--   (1, 1), -- sport: football
--   (1, 2), -- sport: baseball
--   (5, 1), -- sport: football
--   (5, 2), -- sport: baseball
--   (5, 10), -- music: classical music
--   (6, 12), -- music: jazz
--   (6, 17), -- films: action
--   (6, 20), -- films: sci-fi
--   (6, 9), -- music: hip-hop
--   (6, 4), -- sport: ski
--   (6, 21), -- sport: ski
--   (13, 1), -- sport: ski
--   (13, 2), -- sport: ski
--   (13, 4), -- sport: ski
--   (13, 5); -- sport: ski

--- event attendees
CREATE TABLE event_attendees (
  event_id BIGINT,
  user_id  BIGINT,
  PRIMARY KEY (event_id, user_id)
);

CREATE TABLE debug (
  is_updated BOOLEAN PRIMARY KEY
);

INSERT INTO debug VALUES (FALSE);
-- INSERT INTO
--   event_attendees
-- VALUES
--   (1, 1),
--   (2, 1),
--   (5, 1),
--   (7, 1),
--
--   (1, 3),
--   (2, 2),
--   (2, 3),
--   (3, 3),
--   (4, 3);


