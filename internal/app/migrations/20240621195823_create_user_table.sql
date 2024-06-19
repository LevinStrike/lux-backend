-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
   id serial NOT NULL PRIMARY KEY,
   email TEXT NOT NULL UNIQUE,
   password TEXT NOT NULL
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
