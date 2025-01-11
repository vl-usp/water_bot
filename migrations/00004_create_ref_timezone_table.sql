-- +goose Up
-- +goose StatementBegin
create table if not exists ref_timezone (
	id serial primary key,
	name text,
	cities text,
    utc_offset smallint
);

INSERT INTO ref_timezone (name, cities, utc_offset) VALUES
    ('UTC-12', 'Бейкер-Айленд, Хауленд-Айленд', -12),
    ('UTC-11', 'Мидуэй, Ниуэ', -11),
    ('UTC-10', 'Гонолулу, Таити, Раиваваэ', -10),
    ('UTC-9', 'Анкоридж, Густавус, Джуно', -9),
    ('UTC-8', 'Лос-Анджелес, Сан-Франциско, Ванкувер', -8),
    ('UTC-7', 'Денвер, Феникс, Эдмонтон', -7),
    ('UTC-6', 'Чикаго, Мехико, Гватемала-Сити', -6),
    ('UTC-5', 'Нью-Йорк, Торонто, Богота', -5),
    ('UTC-4', 'Сантьяго, Галифакс, Каракас', -4),
    ('UTC-3', 'Буэнос-Айрес, Бразилиа, Монтевидео', -3),
    ('UTC-2', 'Южная Георгия и Южные Сандвичевы острова', -2),
    ('UTC-1', 'Азорские острова, Кабо-Верде', -1),
    ('UTC+0', 'Лондон, Лиссабон, Рейкьявик', 0),
    ('UTC+1', 'Париж, Берлин, Рим', 1),
    ('UTC+2', 'Афины, Киев, Йоханнесбург', 2),
    ('UTC+3', 'Москва, Дубай, Найроби', 3),
    ('UTC+4', 'Баку, Ереван, Самара', 4),
    ('UTC+5', 'Ташкент, Карачи, Екатеринбург', 5),
    ('UTC+6', 'Алма-Ата, Дакка, Бишкек', 6),
    ('UTC+7', 'Джакарта, Бангкок, Ханой', 7),
    ('UTC+8', 'Пекин, Сингапур, Гонконг', 8),
    ('UTC+9', 'Токио, Сеул, Якутск', 9),
    ('UTC+10', 'Сидней, Владивосток, Порт-Морсби', 10),
    ('UTC+11', 'Сувва, Магадан, Гонлонг (Соломоновы острова)', 11),
    ('UTC+12', 'Окленд, Сува, Фунафути', 12),
    ('UTC+13', 'Нукуалофа, Апиа', 13),
    ('UTC+14', 'Киритимати', 14);

create unique index idx_ref_timezone_utc_offset on ref_timezone (name);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ref_timezone;
-- +goose StatementEnd
