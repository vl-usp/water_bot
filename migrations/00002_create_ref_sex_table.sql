-- +goose Up
-- +goose StatementBegin
create table if not exists ref_sex (
	id serial primary key,
	name text
);

insert into ref_sex (name) values ('male'), ('female');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ref_sex;
-- +goose StatementEnd
