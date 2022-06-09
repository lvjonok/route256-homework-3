-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_orders
(
    order_id   SERIAL PRIMARY KEY,
    user_id    INT       NOT NULL,
    status     TEXT      NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE order_items
(
    order_id   INT REFERENCES user_orders (order_id),
    product_id INT       NOT NULL,
    quantity   INT       NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
