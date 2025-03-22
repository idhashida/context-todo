create table tasks (
	id serial primary key,
    list_id int default 1 not null,
    title varchar(500) not null,
    "desc" varchar(500),
    context varchar(2000),
    scenario varchar(2000),
    criterion varchar(2000),
    status_id int default 1 not null,
    deadline timestamp default current_timestamp,
    sublist_id int not null,
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp,
    priority_id int default 1 not null,
    is_del int default 0 not null
)