-- +goose Up

create table todo(
    id UUID primary key,
    name varchar(100),
    owner UUID not null,
    status int not null default 0,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

alter table todo add constraint fk_todo_user foreign key (owner) references users(id) on delete cascade;


-- +goose Down
alter table todo drop constraint fk_todo_user;
drop table todo;