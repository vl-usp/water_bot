-- +goose Up
-- +goose StatementBegin
create table if not exists user_data (
	id serial primary key,
	user_id integer not null unique,
	sex_id integer,
	weight integer,
	physical_activity_id integer,
	climate_id integer,
	water_goal integer not null,
	created_at timestamp not null default 'now()',
	updated_at timestamp,
	foreign key (user_id) references users(id),
	foreign key (sex_id) references ref_sex(id),
	foreign key (physical_activity_id) references ref_physical_activity(id),
	foreign key (climate_id) references ref_climate(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_data;
-- +goose StatementEnd
