-- +goose Up
-- +goose StatementBegin
create table if not exists ref_sex (
	id serial primary key,
	key text,
	name text,
	water_coef float
);

insert into ref_sex (key, name, water_coef) values
	('male', 'Мужчина', 1.0),
	('female', 'Женщина', 0.9);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ref_sex;
-- +goose StatementEnd
