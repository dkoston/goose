
-- +goose Up
-- +goose StatementBegin
CREATE TABLE sql (
    id int NOT NULL PRIMARY KEY
);

INSERT INTO sql (id) VALUES(1);
-- +goose StatementEnd

-- +goose Down
DROP TABLE post;
