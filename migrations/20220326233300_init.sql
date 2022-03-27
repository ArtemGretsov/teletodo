-- +goose Up
-- +goose StatementBegin
create table users(
    id         bigserial not null constraint users_pkey primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name       varchar not null,
    uuid       uuid not null
);

create index idx_users_deleted_at on users (deleted_at);
create index idx_users_uuid on users (uuid);

create type user_auth_type as enum ('telegram');

create table user_authentications(
    id bigserial not null constraint user_authentications_pkey primary key,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,
    external_id varchar not null,
    user_id     bigint not null constraint fk_users_user_authentication references users,
    auth_type   user_auth_type not null
);

create unique index idx_external_id_user_id_auth_type
    on user_authentications (external_id, auth_type, user_id);

create index idx_user_authentications_deleted_at
    on user_authentications (deleted_at);

create table todos(
    id         bigserial not null constraint todos_pkey primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name       text not null,
    author_id  bigint constraint fk_users_todo references users
);

create index idx_todos_deleted_at
    on todos (deleted_at);

create unique index idx_todos_name
    on todos (name)
    where (deleted_at IS NULL);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todos;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS user_authentications;
-- +goose StatementEnd
