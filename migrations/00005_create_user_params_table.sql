-- +goose Up
-- +goose StatementBegin
create table if not exists user_params (
	id serial primary key,
	sex_id smallint,
	physical_activity_id smallint,
	climate_id smallint,
	timezone_id smallint not null,
	weight smallint,
	water_goal smallint not null,
	created_at timestamp not null default 'now()',
	updated_at timestamp,
	foreign key (sex_id) references ref_sex(id),
	foreign key (physical_activity_id) references ref_physical_activity(id),
	foreign key (climate_id) references ref_climate(id),
	foreign key (timezone_id) references ref_timezone(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_params;
-- +goose StatementEnd
