-- +goose Up
-- +goose StatementBegin

ALTER TABLE user_orders ADD COLUMN saga_status TEXT NOT NULL DEFAULT 'booked';

CREATE TABLE retries (
    order_id INT NOT NULL,
    last_status TEXT NOT NULL,

    created_at timestamp not null default now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE user_orders DROP COLUMN saga_status;

-- +goose StatementEnd
