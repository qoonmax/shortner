-- +goose Up
-- +goose StatementBegin
CREATE TABLE urls
(
    id    INTEGER PRIMARY KEY,
    alias VARCHAR(255) NOT NULL UNIQUE,
    url   TEXT         NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE urls;
-- +goose StatementEnd