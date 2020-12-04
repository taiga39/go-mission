
-- +migrate Up
create table users (
id serial PRIMARY KEY,
name text NOT NULL,
token text NOT NULL);
-- +migrate Down
drop table users;