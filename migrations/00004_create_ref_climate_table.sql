-- +goose Up
-- +goose StatementBegin
create table if not exists ref_climate (
	id serial primary key,
	name text
);

insert into ref_climate (name) values ('cold'), ('temperate'), ('warm'), ('hot');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ref_climate;
-- +goose StatementEnd
