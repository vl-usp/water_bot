-- +goose Up
-- +goose StatementBegin
create table if not exists users (
	id integer primary key,
	first_name text,
	last_name text,
	username text not null,
	language_code text,
	created_at timestamp not null default 'now()'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
