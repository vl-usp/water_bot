-- +goose Up
-- +goose StatementBegin
create table if not exists ref_physical_activity (
	id serial primary key,
	key text,
	name text,
	water_coef float
);

insert into ref_physical_activity (key, name, water_coef) values
	('low', 'Низкий', 1),
	('moderate', 'Средний', 1.1),
	('high', 'Высокий', 1.2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ref_physical_activity;
-- +goose StatementEnd
