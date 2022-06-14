-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION postgres_fdw;
CREATE TABLE reviews_even
(
    id         SERIAL,
    product_id INT NOT NULL,
    text       TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
