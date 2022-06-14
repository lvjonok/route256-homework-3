-- +goose Up
-- +goose StatementBegin

-- we want to shard reviews by even and odd product ids

-- rename old reviews table
alter table review rename to review_old;

-- create the same table, but with partition specified
CREATE TABLE review
(
    id         SERIAL,
    product_id INT NOT NULL,
    text       TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
) PARTITION BY HASH (id);

-- create shard server
CREATE EXTENSION IF NOT EXISTS postgres_fdw;
CREATE SERVER even_reviews FOREIGN DATA WRAPPER postgres_fdw
    OPTIONS (host 'db-marketplace-even', port '5432', dbname 'root'); 
CREATE USER MAPPING FOR root SERVER even_reviews OPTIONS (user 'root', password 'root');

-- create foreign table of shard
CREATE FOREIGN TABLE IF NOT EXISTS reviews_even 
    PARTITION OF review FOR VALUES WITH (MODULUS 2, REMAINDER 0) SERVER even_reviews;

-- create partition in the current shard
CREATE TABLE reviews_odd PARTITION OF review FOR VALUES WITH (MODULUS 2, REMAINDER 1);

-- insert the data back
INSERT INTO review(id, product_id, text, created_at) TABLE review_old;
DROP TABLE review_old;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
