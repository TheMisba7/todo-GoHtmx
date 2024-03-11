-- +goose Up
create table task(
    id UUID primary key,
    name varchar(100),
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    status int not null default 0,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    todo_id UUID not null
);

alter table task add constraint fk_task_todo foreign key (todo_id) references todo(id) on delete cascade;

-- +goose Down

alter table task drop constraint fk_task_todo;
drop table task;