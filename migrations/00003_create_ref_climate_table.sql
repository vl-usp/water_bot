-- +goose Up
-- +goose StatementBegin
create table if not exists ref_climate (
	id serial primary key,
	key text,
	name text,
	water_coef float
);

insert into ref_climate (key, name, water_coef) values
	('cold', 'Холодный', 1),
	('temperate', 'Умеренный', 1.1),
	('warm', 'Теплый', 1.2),
	('hot', 'Жаркий', 1.3);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ref_climate;
-- +goose StatementEnd
