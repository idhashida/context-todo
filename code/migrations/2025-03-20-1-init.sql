create table users (
    id serial primary key,
    username varchar(200) not null,
    email varchar(200) not null,
    password_hash varchar(8000) not null,
    created_at timestamp default current_timestamp not null,
    updated_at timestamp
)