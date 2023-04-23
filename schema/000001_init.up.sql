CREATE TABLE users (
    id serial primary key,
    username varchar(128) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    email varchar(128) NOT NULL UNIQUE,
    verified bool NOT NULL default false,
    code varchar(64) NOT NULL
);

CREATE TABLE notes (
    id serial primary key,
    title varchar(128) NOT NULL,
    text varchar(1024) NOT NULL,
    important boolean NOT NULL,
    tags varchar(512),
    user_id int references users (id) on delete cascade
);
