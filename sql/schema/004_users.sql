-- +goose Up

CREATE TABLE users (
  user_id     TEXT PRIMARY KEY,
  country     TEXT NOT NULL,
  age         INT NOT NULL,
  device      TEXT NOT NULL,
  language    TEXT NOT NULL,
  interests   TEXT[] NOT NULL DEFAULT '{}'
);

-- +goose Down
DROP TABLE users;
