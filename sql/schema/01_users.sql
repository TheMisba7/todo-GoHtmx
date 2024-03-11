-- +goose Up
create table users(
    id UUID primary key,
    username varchar(60) not null unique ,
    password TEXT
);

-- +goose Down
drop table users;