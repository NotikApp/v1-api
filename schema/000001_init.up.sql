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
    user_id int references users (id) on delete cascade,
    colour varchar(64) NOT NULL default 'white',
    created_at timestamp NOT NULL
);

CREATE TABLE users_directories (
    id serial primary key,
    user_id int references users (id) on delete cascade,
    directory_id int references directories (id) on delete cascade,
);

CREATE TABLE directories (
    id serial primary key,
    name varchar(128) NOT NULL,
    description varchar(128) NOT NULL
);

CREATE TABLE notes_directories (
    id serial primary key,
    note_id int references notes (id) on delete cascade,
    directory_id int references directories (id) on delete cascade,
);
