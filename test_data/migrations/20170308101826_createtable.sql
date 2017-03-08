
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE testing (
  id BIGINT PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

