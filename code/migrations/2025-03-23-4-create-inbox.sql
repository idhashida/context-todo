create table inbox (
	id serial primary key,
    user_id int not null,
    task_id int not null
)