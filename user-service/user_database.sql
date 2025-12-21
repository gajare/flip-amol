create table users (
    id serial primary key,
    username varchar(50) unique not null,
    email varchar(100) unique not null,
    password_hash varchar(255) not null,
    Role varchar(50) default 'admin',
    created_at timestamp default current_timestamp
);