CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    time_created timestamp NOT NULL,
    time_updated timestamp NOT NULL,
    username varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL,
    name varchar(255),
    surname varchar(255),
    birthday date,
    email varchar(255),
    phone varchar(255)
);
