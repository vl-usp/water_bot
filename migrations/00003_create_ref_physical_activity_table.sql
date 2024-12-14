-- +goose Up
-- +goose StatementBegin
create table if not exists ref_physical_activity (
	id serial primary key,
	name text
);

insert into ref_physical_activity (name) values ('low'), ('moderate'), ('high');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ref_physical_activity;
-- +goose StatementEnd
