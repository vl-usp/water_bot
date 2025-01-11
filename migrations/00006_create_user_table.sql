-- +goose Up
-- +goose StatementBegin
create table if not exists users (
	id bigint primary key,
	first_name text,
	last_name text,
	username text not null,
	language_code text,
	params_id bigint,
	created_at timestamp not null default 'now()',
	foreign key (params_id) references user_params(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
