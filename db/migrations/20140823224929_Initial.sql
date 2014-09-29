-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users
(
  id bigserial NOT NULL,
  name text,
  email text,
  hash text,
  CONSTRAINT users_pkey PRIMARY KEY (id),
  CONSTRAINT users_email UNIQUE (email)
);

CREATE TABLE sessions
(
  id serial NOT NULL,
  user_id bigint,
  api_key character varying(128),
  CONSTRAINT sessions_pk PRIMARY KEY (id)
);

CREATE TABLE products
(
  id bigserial NOT NULL,
  name text,
  description text,
  price double precision,
  category_ids bigint[] NOT NULL DEFAULT '{}',
  CONSTRAINT product_pkey PRIMARY KEY (id)
);

CREATE TABLE categories
(
  id          bigserial     NOT NULL,
  parent_id   bigint,
  name        varchar(100)  NOT NULL,
  fq_name     varchar(200)  NOT NULL DEFAULT '',
  CONSTRAINT categories_pkey PRIMARY KEY (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE categories
DROP TABLE products;
DROP TABLE sessions;
DROP TABLE users;
