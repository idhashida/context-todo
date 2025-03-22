create table lists (
  id serial primary key,
  name varchar(200) not null,
  user_id int not null,
  color varchar(50) not null,
  created_at timestamp default current_timestamp not null,
  updated_at timestamp
)