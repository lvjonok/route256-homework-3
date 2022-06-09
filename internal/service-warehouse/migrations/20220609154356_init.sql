-- +goose Up
-- +goose StatementBegin
CREATE TABLE entries
(
    id         SERIAL PRIMARY KEY,
    product_id INT  NOT NULL,
    quantity   INT  NOT NULL,
    deleted    BOOL NOT NULL DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
