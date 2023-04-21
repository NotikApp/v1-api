CREATE TABLE users (
    id serial primary key,
    username varchar(128) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    email varchar(128) NOT NULL UNIQUE
);

CREATE TABLE notes (
    id serial primary key,
    title varchar(128) NOT NULL,
    text varchar(1024) NOT NULL,
    important boolean NOT NULL,
    tags json,
    user_id int references users (id) on delete cascade
);