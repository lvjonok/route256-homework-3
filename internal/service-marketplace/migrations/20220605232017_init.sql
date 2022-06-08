-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    new.updated_at = NOW();
    RETURN new;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE product
(
    id         SERIAL PRIMARY KEY,
    name       TEXT      NOT NULL DEFAULT '',
    "desc"     TEXT      NOT NULL DEFAULT '',

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE review
(
    id         SERIAL PRIMARY KEY,
    product_id SERIAL REFERENCES product (id),
    text       TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE cart
(
    id         SERIAL    PRIMARY KEY,
    user_id    INT       NOT NULL,
    product_id INT REFERENCES product (id),
    quantity   INT       NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted BOOL NOT NULL DEFAULT FALSE
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table cart;
drop table review;
drop table product;
-- +goose StatementEnd
