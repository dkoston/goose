-- NO TRANSACTIONS --
-- +goose Up
-- +goose StatementBegin
CREATE TABLE post (
    id int NOT NULL,
    title text,
    body text,
    PRIMARY KEY(id)
);

INSERT INTO post (id, title, body) VALUES(1, 'test post', 'this is a test post');
-- +goose StatementEnd

-- +goose Down
DROP TABLE post;